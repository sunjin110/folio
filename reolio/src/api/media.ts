import { CreateMediumOutput } from "./response/media";
import { CreateMediumInput } from "./request/media";
import { ErrorResponse } from "./api";

export async function createMedium(
  fileName: string,
): Promise<CreateMediumOutput | ErrorResponse> {
  const input: CreateMediumInput = {
    file_name: fileName,
  };

  const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + "/media", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(input),
    credentials: "include",
  });

  if (!resp.ok) {
    return {
      type: "error",
      message: await resp.text(),
      status: resp.status,
    };
  }

  const output: CreateMediumOutput = await resp.json();
  return output;
}
