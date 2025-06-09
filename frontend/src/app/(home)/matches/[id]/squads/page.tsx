export default async function Squads({ params }: { params: Promise<{ id: string }> }) {
  console.log(params);
  return <div className="flex flex-col gap-4">Squads</div>;
}
