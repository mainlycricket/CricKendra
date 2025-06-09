export default async function Overs({ params }: { params: Promise<{ id: string }> }) {
  console.log(params);
  return <div className="flex flex-col gap-4">Overs</div>;
}
