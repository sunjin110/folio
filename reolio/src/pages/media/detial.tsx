import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { Medium } from "@/domain/model/media";
import { AuthError, InternalError } from "@/error/error";
import { MediaUsecase } from "@/usecase/media";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export interface MediaDetailProps {
  mediaUsecase: MediaUsecase;
}

export default function MediaDetail(props: MediaDetailProps) {
  const { mediaUsecase } = props;

  const { mediaId } = useParams();

  const [medium, setMedium] = useState<Medium | null>(null);
  const { toast } = useToast();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchMedium = async (id: string) => {
      try {
        const resp = await mediaUsecase.GetMedium(id);
        setMedium(resp);
      } catch (err) {
        if (err instanceof AuthError) {
          toast({
            title: "🔑 Please login again 🔒",
            description: err.message,
          });
          navigate("/login");
          return;
        } else if (err instanceof InternalError) {
          toast({
            title: "🚫 Error 🚫",
            description: err.message,
          });
          return;
        }
        toast({
          title: "Error",
          description: `${err}`,
        });
        console.error(err);
      }
    };

    if (mediaId) {
      fetchMedium(mediaId);
    }
  }, [mediaId, toast, navigate, mediaUsecase]);

  return (
    <Navigation title="Media" sidebarPosition="media">
      <div className="p-2">
        {medium ? (
          <div className="flex flex-col items-center space-y-4">
            <a href={medium.downloadUrl}>
              <img
                key={medium.id}
                src={medium.downloadUrl}
                alt="media"
                className="max-h-64 h-auto shadow-lg rounded"
                onError={(e) =>
                  (e.currentTarget.src = "/image/no_image_square.jpg")
                }
              />
            </a>
            <Button asChild>
              <a href={medium.downloadUrl} download className="inline-block">
                TODO: Download Media
              </a>
            </Button>
          </div>
        ) : (
          <div className="text-center text-gray-700">no image</div>
        )}
      </div>
    </Navigation>
  );
}
