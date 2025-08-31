import { IAllSeries } from "@/lib/types/series.types";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import Link from "next/link";
import { Badge } from "../ui/badge";

export function SingleSeries({ series }: { series: IAllSeries }) {
  return (
    <Link href={`/series/${series.id}`} className="hover:text-blue-500">
      <Card>
        <CardHeader>
          <CardTitle>
            {series.name}, {series.season} {!series.is_male && <Badge className="secondary">Women</Badge>}
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div>
            <span className="font-medium">Teams: </span>
            <span>{series?.teams?.map((team) => team.name)?.join(", ")}</span>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}
