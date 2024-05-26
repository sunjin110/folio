import { MediumSummary } from "../model/media";

export interface MediaRepository {
    CreateMedium(fileName: string): Promise<string>;
    UploadFile(url: string, file: File): void;
    FindMediumSummaries(offset: number, limit: number): Promise<FindMediumSummariesOutput>;
}

export interface FindMediumSummariesOutput {
    totalCount: number;
    summaries: MediumSummary[];
}