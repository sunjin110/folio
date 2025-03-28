use anyhow::Ok;
use aws_lambda_events::cloudwatch_events::macie::FeatureInfo;

use crate::domain::repository::media::MediaRepo;
use aws_sdk_s3::Client;


pub struct MediaRepoImpl {
    s3_client: Client,
    bucket_name: String,
}

pub fn new_media_repo(s3_client: Client, bucket_name: String) -> impl MediaRepo {
    MediaRepoImpl{
        s3_client,
        bucket_name,
    }
}


impl MediaRepo for MediaRepoImpl {
    // https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/rustv1/examples/s3/src/bin/get-object.rs#L20
    // ↑参考
    async fn get_object(&self, key: &str) -> anyhow::Result<()>{

        let mut object = self.s3_client.get_object().bucket(&self.bucket_name).key(key).send().await?;

        let mut byte_count = 0_usize;

        while let Some(bytes) = object.body.try_next().await? {

        }

        Ok(())
    }
}
