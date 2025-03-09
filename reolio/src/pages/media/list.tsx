import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { MediumSummary } from "@/domain/model/media";
import { AuthError, InternalError } from "@/error/error";
// import { useImageNew } from "@/hooks/useImageNew";
import { MediaUsecase } from "@/usecase/media";
import { useCallback, useEffect, useRef, useState } from "react";
import { useDropzone } from "react-dropzone";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

export interface MediaProps {
  mediaUsecase: MediaUsecase;
}

const SUMMARY_PAGING_LIMIT = 10;

function useUniqueSummaries(
  initialData: MediumSummary[],
  initialIdMap?: Map<string, boolean>,
) {
  const [summaries, setSummaries] = useState<MediumSummary[]>(initialData);
  const idMapRef = useRef<Map<string, boolean>>(initialIdMap || new Map());

  const addSummaries = useCallback((newData: MediumSummary[]) => {
    const uniqueData = newData.filter((item) => {
      if (idMapRef.current.has(item.id)) return false;
      idMapRef.current.set(item.id, true);
      return true;
    });
    setSummaries((prev) => [...prev, ...uniqueData]);
  }, []);

  return { summaries, addSummaries };
}

export default function Media(props: MediaProps) {
  const { mediaUsecase } = props;

  const { toast } = useToast();
  const navigate = useNavigate();

  const observer = useRef<IntersectionObserver>(null); // scroll
  const [page, setPage] = useState(0);
  const [hasMore, setHasMore] = useState(true);

  const { summaries, addSummaries } = useUniqueSummaries([]);

  const [previeImagePaths, setPreviewImagePaths] = useState<string[]>([]);
  const [uploadImagePaths, setUploadImagePathes] = useState<string[]>([]);

  const [files, setFiles] = useState<File[]>();
  const onDrop = (acceptedFiles: File[]) => {
    setFiles(acceptedFiles);
    const dataUrls = acceptedFiles.map((file) => URL.createObjectURL(file));
    setPreviewImagePaths(dataUrls);
  };

  const { getRootProps, getInputProps } = useDropzone({
    onDrop: onDrop,
    noClick: true,
  });

  const lastMediaElementRef = useCallback(
    (node: Element | null) => {
      if (observer.current) observer.current.disconnect();
      observer.current = new IntersectionObserver((entries) => {
        if (entries[0].isIntersecting && hasMore) {
          setPage((prevPageNumber) => prevPageNumber + 1);
        }
      });
      if (node) observer.current.observe(node);
    },
    [hasMore],
  );

  const handleUploadButton = useCallback(async () => {
    if (!files || files.length === 0) {
      toast({
        title: "📁 Please select files 📁",
      });
      return;
    }

    try {
      await mediaUsecase.UploadFiles(files);
    } catch (err) {
      console.error(err);
      if (err instanceof AuthError) {
        toast({
          title: "AuthError",
          description: err.message,
        });
        return;
      } else if (err instanceof InternalError) {
        toast({
          title: "InternalError",
          description: err.message,
        });
        return;
      }
      toast({
        title: "Error",
        description: `${err}`,
      });
    }

    setUploadImagePathes((prev) => [...previeImagePaths, ...prev]);
    setPreviewImagePaths([]);
    setFiles([]);
  }, [files, mediaUsecase, previeImagePaths, toast]);

  useEffect(() => {
    const loadMoreMedia = async () => {
      try {
        const output = await mediaUsecase.FindMedia(
          page * SUMMARY_PAGING_LIMIT,
          SUMMARY_PAGING_LIMIT,
        );
        if (output.summaries.length === 0) {
          setHasMore(false);
          return;
        }

        addSummaries(output.summaries);
      } catch (err) {
        if (err instanceof AuthError) {
          toast({
            title: "🔑 Please login 🔒",
            description: err.message,
          });
          navigate("/login");
          return;
        }
        toast({
          title: "🚫 Error 🚫",
          description: `${err}`,
        });
        return;
      }
    };
    loadMoreMedia();
  }, [page, mediaUsecase, addSummaries, toast, navigate]);

  useEffect(() => {
    if (files && files.length > 0) {
      handleUploadButton();
    }
  }, [files, handleUploadButton]);

  const filesInputRef = useRef<HTMLInputElement>(null);

  const clickFileSelect = () => {
    if (filesInputRef.current) {
      filesInputRef.current.click();
    }
  };

  const onFileInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (!event.target.files) {
      return;
    }
    let targetFiles: File[] = [];
    for (let i = 0; i < event.target.files.length; i++) {
      const f = event.target.files.item(i);
      if (!f) {
        continue;
      }
      targetFiles.push(f);
    }
    setFiles(targetFiles);
    const dataUrls = targetFiles.map((file) => URL.createObjectURL(file));
    setPreviewImagePaths(dataUrls);
  };

  return (
    <Navigation title="Media" sidebarPosition="media">
      <div {...getRootProps({ className: "dropzone" })}>
        <input {...getInputProps()} />

        <div className="mb-4">
          <Button onClick={clickFileSelect}>📁 Upload files</Button>
          <input
            type="file"
            hidden
            onChange={onFileInputChange}
            ref={filesInputRef}
          />
        </div>

        <div className="grid grid-cols-3 md:grid-cols-5 lg:grid-cols-6 xl:grid-cols-8 gap-4">
          {previeImagePaths &&
            previeImagePaths.map((path) => {
              return (
                <div
                  key={path}
                  className="w-full aspect-square overflow-hidden relative group"
                >
                  <img
                    src={path}
                    key={path}
                    alt="preview"
                    className="w-full h-full object-cover object-center transition-transform duration-300 ease-in-out transform group-hover:scale-110"
                  />
                  <div className="absolute inset-0 bg-white bg-opacity-50"></div>
                </div>
              );
            })}

          {uploadImagePaths &&
            uploadImagePaths.map((path) => {
              return (
                <div
                  key={path}
                  className="w-full aspect-square overflow-hidden relative group"
                >
                  <img
                    src={path}
                    key={path}
                    alt="doned"
                    className="w-full h-full object-cover object-center transition-transform duration-300 ease-in-out transform group-hover:scale-110"
                  />
                </div>
              );
            })}

          {summaries &&
            summaries.map((summary, index) => {
              const isLastElement = index === summaries.length - 1;

              return (
                <div
                  key={summary.id}
                  ref={isLastElement ? lastMediaElementRef : null}
                  className="w-full aspect-square overflow-hidden relative group"
                >
                  <Link to={`/media/${summary.id}`}>
                    <img
                      src={summary.thumbnailUrl}
                      alt="Thumbnail"
                      className="w-full h-full object-cover object-center transition-transform duration-300 ease-in-out transform group-hover:scale-110"
                      onError={(e) =>
                        (e.currentTarget.src = "/image/no_image_square.jpg")
                      }
                    />
                  </Link>
                </div>
              );
            })}
        </div>
      </div>
    </Navigation>
  );
}
