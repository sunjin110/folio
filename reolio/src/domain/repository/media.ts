import { Medium, MediumSummary } from "../model/media";

export interface MediaRepository {
  CreateMedium(fileName: string): Promise<string>;
  UploadFile(url: string, file: File): void;
  FindMediumSummaries(
    offset: number,
    limit: number,
  ): Promise<FindMediumSummariesOutput>;
  GetMedium(id: string): Promise<Medium>;
}

export interface FindMediumSummariesOutput {
  totalCount: number;
  summaries: MediumSummary[];
}
