use image::DynamicImage;

pub trait ImageRepo {
    fn resize_image(&self, img: DynamicImage, width: u32, height: u32) -> anyhow::Result<DynamicImage>;
    fn trim_center(&self, img: DynamicImage, width: u32, height: u32) -> anyhow::Result<DynamicImage>;
}
