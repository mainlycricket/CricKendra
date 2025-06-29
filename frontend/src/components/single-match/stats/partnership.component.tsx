import { IPartnershipStats } from "@/lib/types/single-match.types";

export function Partnerships({ partnerships }: { partnerships: IPartnershipStats[] }) {
  partnerships.sort((a, b) => a.for_wicket - b.for_wicket);

  const highestPartnership = Math.max(...partnerships.map((partnership) => partnership.partnership_runs));

  return (
    <div className="mt-4 px-2 text-sm flex flex-col gap-4 md:w-full">
      {partnerships.map((partnership) => {
        const startOver = Math.floor(partnership.start_ball_number);
        const endOver = Math.floor(partnership.end_ball_number);

        const balls =
          (endOver - startOver) * 6 -
          Math.round((partnership.start_ball_number - startOver) * 10 - 1) +
          Math.round((partnership.end_ball_number - endOver) * 10);

        return (
          <div key={`${partnership.for_wicket}_${partnership.batter1_id}_${partnership.batter2_id}`}>
            <div className="flex">
              <p className="w-1/3">{partnership.batter1_name}</p>
              <p className="w-1/3 text-center">
                {partnership.partnership_runs} ({balls})
              </p>
              <p className="w-1/3 text-right">{partnership.batter2_name}</p>
            </div>

            <div className="flex justify-between w-full">
              <p className="w-24">
                {partnership.batter1_runs} ({partnership.batter1_balls})
              </p>
              <div
                className="flex h-2"
                style={{
                  width: `${(partnership.partnership_runs * 100) / highestPartnership}%`,
                }}
              >
                <div
                  style={{
                    width: `${(partnership.batter1_runs * 100) / partnership.partnership_runs}%`,
                  }}
                  className="bg-[#B3181E] rounded-l-lg"
                ></div>
                <div
                  style={{
                    width: `${(partnership.batter2_runs * 100) / partnership.partnership_runs}%`,
                  }}
                  className="bg-[#B3181E] opacity-50 rounded-r-lg"
                ></div>
              </div>
              <p className="w-24 text-right">
                {partnership.batter2_runs} ({partnership.batter2_balls})
              </p>
            </div>
          </div>
        );
      })}
    </div>
  );
}
