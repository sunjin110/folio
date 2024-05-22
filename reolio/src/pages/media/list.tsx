import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { MediumSummary } from "@/domain/model/media";
import { AuthError, InternalError } from "@/error/error";
// import { useImageNew } from "@/hooks/useImageNew";
import { MediaUsecase } from "@/usecase/media";
import { useCallback, useEffect, useRef, useState } from "react";
import { useDropzone } from "react-dropzone";
import { useNavigate } from "react-router-dom";

export interface MediaProps {
    mediaUsecase: MediaUsecase
}

function useUniqueSummaries(initialData: MediumSummary[], initialIdMap?: Map<string, boolean>) {
    const [summaries, setSummaries] = useState<MediumSummary[]>(initialData);
    const idMapRef = useRef<Map<string, boolean>>(initialIdMap || new Map());

    const addSummaries = useCallback((newData: MediumSummary[]) => {
        const uniqueData = newData.filter(item => {
            if (idMapRef.current.has(item.id)) return false;
            idMapRef.current.set(item.id, true);
            return true
        });
        setSummaries(prev => [...prev, ...uniqueData]);
    }, []);

    return { summaries, addSummaries };
}

export default function Media(props: MediaProps) {
    const { mediaUsecase } = props;
    
    const { toast } = useToast();
    const navigate = useNavigate();

    const observer = useRef<IntersectionObserver>(); // scroll
    const [page, setPage] = useState(0);
    const [hasMore, setHasMore] = useState(true);

    const { summaries, addSummaries } = useUniqueSummaries([]);

    // const [previeImagePaths, setPreviewImagePaths] = useState<string[]>();
    const [files, setFiles] = useState<File[]>();
    const onDrop = (acceptedFiles: File[]) => {
        setFiles(acceptedFiles);
        // const dataUrls = acceptedFiles.map((file) => URL.createObjectURL(file));
        // setPreviewImagePaths(dataUrls);
    }

    const { getRootProps, getInputProps } = useDropzone({ 
        onDrop: onDrop,
        noClick: true,
     });

    const lastMediaElementRef = useCallback((node: Element | null) => {
        if (observer.current) observer.current.disconnect();
        observer.current = new IntersectionObserver(entries => {
            if (entries[0].isIntersecting && hasMore) {
                setPage(prevPageNumber => prevPageNumber + 1);
            }
        });
        if (node) observer.current.observe(node);
    }, [hasMore]);

    const handleUploadButton = async () => {
        if (!files || files.length === 0) {
            toast({
                title: 'select files'
            });
            return;
        }

        try {
            await mediaUsecase.UploadFiles(files);
        } catch(err) {
            console.error(err);
            if (err instanceof AuthError) {
                toast({
                    title: 'AuthError',
                    description: err.message,
                });
                return;
            } else if (err instanceof InternalError) {
                toast({
                    title: 'InternalError',
                    description: err.message,
                });
                return;
            }
            toast({
                title: 'Error',
                description: `${err}`,
            })
        }

    };

    useEffect(() => {
        const loadMoreMedia = async () => {
            try {
                console.log("追加ロードが走ったよ");
                const output = await mediaUsecase.FindMedia(page * 3, 3);
                if (output.summaries.length === 0) {
                    console.log("もうないよ!");
                    setHasMore(false);
                    return;
                }

                addSummaries(output.summaries);
            } catch (err) {
                if (err instanceof AuthError) {
                    toast({
                        title: 'Please login',
                        description: err.message,
                    });
                    navigate("/login");
                    return;
                }
                toast({
                    title: 'Error',
                    description: `${err}`
                })
                return;
            }
        }
        loadMoreMedia();
    }, [page, mediaUsecase, addSummaries, toast, navigate]);


    return (
        <Navigation title="Media" sidebarPosition="media">
            <div {...getRootProps({className: "dropzone"})}>
                <input {...getInputProps()} />

                <div className="mb-2">
                <Button onClick={handleUploadButton}>Upload</Button>
                </div>
                
                <div className="grid grid-cols-3 md:grid-cols-5 lg:grid-cols-6 xl:grid-cols-8 gap-4">
                {summaries && (
                    summaries.map((summary, index) => {
                        const isLastElement = index === summaries.length - 1;

                        return (
                            <div 
                            key={summary.id} 
                            ref={isLastElement ? lastMediaElementRef : null}
                            className="w-full aspect-square overflow-hidden relative group"
                            style={{ backgroundImage: `url(/image/no_image_square.jpg)` }}
                        >
                            <img 
                                src={summary.thumbnailUrl} 
                                alt="Thumbnail"
                                className="w-full h-full object-cover object-center transition-transform duration-300 ease-in-out transform group-hover:scale-110"
                                onError={(e) => e.currentTarget.src = '/image/no_image_square.jpg'}
                            />
                        </div>
                        );
                    })
                )}
                </div>
            </div>
        </Navigation>
    )
}
