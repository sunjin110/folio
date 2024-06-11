import { Medium, MediumSummary } from "@/domain/model/media";
import {
  FindMediumSummariesOutput,
  MediaRepository,
} from "@/domain/repository/media";
import { AuthError, InternalError } from "@/error/error";
import {
  GolioApi,
  MediaGet200Response,
  MediaMediumIdGet200Response,
  MediaPost200Response,
  ResponseError,
} from "@/generate/schema/http";

export function NewMediaRepository(golioApi: GolioApi): MediaRepository {
  return new media(golioApi);
}

class media implements MediaRepository {
  private golioApi: GolioApi;

  constructor(golioApi: GolioApi) {
    this.golioApi = golioApi;
  }
  async GetMedium(id: string): Promise<Medium> {
    let resp: MediaMediumIdGet200Response = {
      mediumId: "",
      thumbnailUrl: "",
      downloadUrl: "",
      fileType: "",
    };

    try {
      resp = await this.golioApi.mediaMediumIdGet({
        mediumId: id,
      });
    } catch (err) {
      if (err instanceof ResponseError) {
        if (err.response.status === 401 || err.response.status === 403) {
          throw new AuthError("failed get media", err);
        }
        throw new InternalError("failed get media", err);
      }
      throw err;
    }

    return {
      id: resp.mediumId,
      fileType: resp.fileType,
      downloadUrl: resp.downloadUrl,
    };
  }

  async FindMediumSummaries(
    offset: number,
    limit: number,
  ): Promise<FindMediumSummariesOutput> {
    let resp: MediaGet200Response = {
      totalCount: 0,
      media: [],
    };
    try {
      resp = await this.golioApi.mediaGet({
        offset: offset,
        limit: limit,
      });
    } catch (err) {
      if (err instanceof ResponseError) {
        if (err.response.status === 401 || err.response.status === 403) {
          throw new AuthError("failed find media", err);
        }
        throw new InternalError("failed find media", err);
      }
      throw err;
    }
    let mediumSummaries: MediumSummary[] = [];
    for (let medium of resp.media) {
      mediumSummaries.push({
        id: medium.id,
        fileType: medium.fileType,
        thumbnailUrl: medium.thumbnailUrl,
        createdAt: medium.createdAt,
        updatedAt: medium.updatedAt,
      });
    }
    return {
      totalCount: resp.totalCount,
      summaries: mediumSummaries,
    };
  }

  async CreateMedium(fileName: string): Promise<string> {
    let resp: MediaPost200Response = {
      uploadPresignedUrl: "",
    };
    try {
      resp = await this.golioApi.mediaPost({
        mediaPostRequest: {
          fileName: fileName,
        },
      });
    } catch (err) {
      if (err instanceof ResponseError) {
        if (err.response.status === 401 || err.response.status === 403) {
          throw new AuthError("failed create medium", err);
        }
        throw new InternalError("failed create medium", err);
      }
    }
    return resp.uploadPresignedUrl;
  }

  UploadFile(url: string, file: File): void {
    const xhr = new XMLHttpRequest();
    xhr.open("PUT", url, true);
    xhr.setRequestHeader("Content-Type", file.type);

    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable) {
        const percentCompleted = Math.round((event.loaded * 100) / event.total);
        console.log("progress: ", percentCompleted);
      }
    };

    xhr.onload = () => {
      if (xhr.status === 200) {
        console.info("success file upload", file.name);
        return;
      }
      throw new Error(`failed upload file: ${file.name}`);
    };
    xhr.send(file);
  }
}
