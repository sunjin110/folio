
pub trait MovieRepoTrait {
    fn new() -> Self;
    fn generate_thumbnail_file(input: &str, output: &str, time: i64) -> anyhow::Result<()>;
}
