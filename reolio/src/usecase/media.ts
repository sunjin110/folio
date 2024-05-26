import { MediumSummary } from "@/domain/model/media";
import { MediaRepository } from "@/domain/repository/media";

export interface MediaUsecase {
    // UploadFile ファイルをアップロードする
    UploadFile(url: string, file: File): void;

    // UploadFiles ファイルをアップロードする
    UploadFiles(files: File[]): Promise<void>;

    // FindMedia メディアの一覧を取得する
    FindMedia(offset: number, limit: number): Promise<FindMediaOutput>;
}

export function NewMediaUsecase(mediaRepo: MediaRepository): MediaUsecase {
    return new media(mediaRepo);
}

export interface FindMediaOutput {
    totalCount: number;
    summaries: MediumSummary[];
}

class media implements MediaUsecase {

    private mediaRepo: MediaRepository

    constructor(mediaRepo: MediaRepository) {
        this.mediaRepo = mediaRepo;
    }
    async FindMedia(offset: number, limit: number): Promise<FindMediaOutput> {
        const output = await this.mediaRepo.FindMediumSummaries(offset, limit);
        return {
            totalCount: output.totalCount,
            summaries: output.summaries,
        };
    }

    UploadFile(url: string, file: File): void {
        try {
            this.mediaRepo.UploadFile(url, file);
        } catch (err) {
            throw new Error(`failed upload file. err: ${err}`);
        }
    }

    async UploadFiles(files: File[]): Promise<void> {
        try {
            // TODO 非同期
            for (let file of files) {
                const presignedUrl = await this.mediaRepo.CreateMedium(file.name);
                this.mediaRepo.UploadFile(presignedUrl, file);
            }
        } catch(err) {
            throw err;
        }
    }
}
