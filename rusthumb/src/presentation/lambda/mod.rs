use aws_lambda_events::s3::S3Event;
use aws_sdk_s3::Client;
use lambda_runtime::{run, service_fn,  Error, LambdaEvent};
use log::info;


pub async fn serve() -> Result<(), Error> {
    let config = aws_config::load_from_env().await;

    let s3_client = Client::new(&config);

    run(service_fn(|request: LambdaEvent<S3Event>| {
        function_handler(&s3_client, request)
    })).await
}

async fn function_handler(_s3_client: &Client, evt: LambdaEvent<S3Event>) -> Result<(), Error> {
    info!("bucket name is {:?}", evt.payload.records[0].s3.bucket.name);
    info!("bucket path is {:?}", evt.payload.records[0].s3.object.key);
    Ok(())
}
