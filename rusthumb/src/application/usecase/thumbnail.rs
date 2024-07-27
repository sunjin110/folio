use crate::domain::repository::image::ImageRepo;

pub trait ThumbnailUsecase {
}

pub fn new_thubnail_usecase(_image_repo: Box<dyn ImageRepo + Send + Sync>) -> impl ThumbnailUsecase {
    return ThumbnailUsecaseImpl{
        // image_repo,
    }
}

struct ThumbnailUsecaseImpl {
    // image_repo: Box<dyn ImageRepo + Send + Sync>,
}

impl ThumbnailUsecase for ThumbnailUsecaseImpl {

}
