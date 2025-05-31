
pub trait MediaRepo {
    fn get_object(&self, key: &str) -> impl std::future::Future<Output = anyhow::Result<()>> + Send;
}
