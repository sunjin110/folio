import { Navigation } from "@/components/organisms/navigation";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Table,
  TableBody,
  TableCell,
  TableRow,
  TableHeader,
  TableHead,
} from "@/components/ui/table";
import { useEffect, useState } from "react";

function toIsoStringWithLocalTimezone(date: Date, timezoneOffset: number) {
  const tzo = -timezoneOffset;
  const dif = tzo >= 0 ? "+" : "-";
  const pad = function (num: number) {
    return (num < 10 ? "0" : "") + num;
  };

  return (
    date.getFullYear() +
    "-" +
    pad(date.getMonth() + 1) +
    "-" +
    pad(date.getDate()) +
    "T" +
    pad(date.getHours()) +
    ":" +
    pad(date.getMinutes()) +
    ":" +
    pad(date.getSeconds()) +
    dif +
    pad(Math.floor(Math.abs(tzo) / 60)) +
    ":" +
    pad(Math.abs(tzo) % 60)
  );
}

export default function Time() {
  const [timestamp, setTimestamp] = useState<number | string>("");
  const [dateStr, setDateStr] = useState("");

  const [dateFromTimestamp, setDateFromTimestamp] = useState<
    Date | undefined
  >();

  const [timestampFromDateStr, setTimestampFromDateStr] = useState<
    number | undefined
  >();

  const now = new Date();

  useEffect(() => {
    if (timestamp !== "") {
      setDateFromTimestamp(new Date(Number(timestamp) * 1000));
    }
  }, [timestamp]);

  useEffect(() => {
    const time = Date.parse(dateStr);
    setTimestampFromDateStr(time / 1000);
  }, [dateStr]);

  return (
    <Navigation title="Time" sidebarPosition="">
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>time</TableHead>
              <TableHead>value</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell key={"unix_time_now_1"}>unix_time</TableCell>
              <TableCell key={"unix_time_now_2"}>
                {Math.ceil(Number(now.getTime() / 1000))}
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell key={"utc_now_1"}>UTC</TableCell>
              <TableCell key={"utc_now_2"}>{now.toISOString()}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell key={"local_now_1"}>Local</TableCell>
              <TableCell key={"local_now_2"}>
                {toIsoStringWithLocalTimezone(now, now.getTimezoneOffset())}
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell key={"jst_now_1"}>JST</TableCell>
              <TableCell key={"jst_now_2"}>
                {toIsoStringWithLocalTimezone(now, -9 * 60)}
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <div className="p-5">
        <Label htmlFor="timestamp">Timestamp</Label>
        <Input
          value={timestamp}
          type="number"
          onChange={(event) => setTimestamp(Number(event.target.value))}
          id="timestamp"
        />
        {dateFromTimestamp ? (
          <div className="p-5">
            UTC: {dateFromTimestamp.toISOString()}
            <br />
            Local:
            {toIsoStringWithLocalTimezone(
              dateFromTimestamp,
              dateFromTimestamp.getTimezoneOffset(),
            )}
            <br />
            JST: {toIsoStringWithLocalTimezone(dateFromTimestamp, -9 * 60)}
          </div>
        ) : (
          <></>
        )}
      </div>
      <div className="p-5">
        <Label>Datetime</Label>
        <Input
          value={dateStr}
          type="text"
          onChange={(event) => setDateStr(event.target.value)}
        />
        {timestampFromDateStr ? (
          <div className="p-5">
            timestamp: {Math.ceil(timestampFromDateStr)}
          </div>
        ) : (
          <></>
        )}
      </div>
    </Navigation>
  );
}
