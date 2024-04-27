import { useEffect, useState } from "react";
import { columns, Payment } from "./columns";
import { DataTable } from "./data-table";


export const payments: Payment[] = [
    {
        id: "728ed52f",
        amount: 100,
        status: "pending",
        email: "m@example.com",
      },
      {
        id: "489e1d42",
        amount: 125,
        status: "processing",
        email: "example@gmail.com",
      },
];

async function getData(): Promise<Payment[]> {
    return payments;
} 

export default function Articles() {
    const [data, setData] = useState<Payment[]>([]);
    useEffect(() => {
        const fetch = async () => {
            setData(await getData());
        };
        fetch();

    });
    return (<div className="container mx-auto py-10">
        <DataTable columns={columns} data={data}></DataTable>
    </div>);
}
