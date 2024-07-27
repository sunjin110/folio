use image::{imageops::FilterType, DynamicImage, GenericImageView};

use crate::domain::repository::image::ImageRepo;

pub struct ImageRepoImpl {}

impl ImageRepo for ImageRepoImpl {
    fn resize_image(&self, img: DynamicImage, width: u32, height: u32) -> anyhow::Result<DynamicImage> {
        let resized_img = img.resize(width, height, FilterType::Nearest);
        anyhow::Ok(resized_img)
    }

    fn trim_center(&self, img: DynamicImage, width: u32, height: u32) -> anyhow::Result<DynamicImage> {
        let (origin_width, origin_height) = img.dimensions();
        let start_x = (origin_width / 2) - (width / 2);
        let start_y = (origin_height / 2) - (height / 2);
        let cropped_img = img.crop_imm(start_x, start_y, width, height);
        anyhow::Ok(cropped_img)
    }
}
