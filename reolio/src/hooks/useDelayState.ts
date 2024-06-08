import { useEffect, useState, useRef } from "react";

export const useDelayState = <T>(value: T, delayMs: number) => {
    const [delayValue, setDelayValue] = useState<T>();
    const timer = useRef<NodeJS.Timeout | null>(null);

    // タイマーをクリアするための useEffect は、コンポーネントのクリーンアップ時のみ実行する
    useEffect(() => {
        return () => {
          if (timer.current) clearTimeout(timer.current);
        };
    }, []);

    useEffect(() => {
        if (timer.current) {
            clearTimeout(timer.current);
        }

        timer.current = setTimeout(() => {
            setDelayValue(value);
        }, delayMs);

        // この useEffect のクリーンアップ関数でタイマーをクリア
        return () => {
            if (timer.current) clearTimeout(timer.current);
        };

    }, [value, delayMs]); // timer は依存関係リストから除外し、useRefを使用

    return [delayValue];
};
