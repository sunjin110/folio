use rusthumb::presentation::lambda;
use lambda_runtime::Error;

#[tokio::main]
async fn main() -> Result<(), Error> {
    lambda::serve().await
}
