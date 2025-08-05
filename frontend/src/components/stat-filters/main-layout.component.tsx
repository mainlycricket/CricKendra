"use client";

import { IStatsFilters } from "@/lib/types/filters-stats.types";
import { CommonFilters } from "./common-filters.component";
import Link from "next/link";
import { Button } from "../ui/button";
import { Dispatch, FormEvent, SetStateAction, useEffect, useState } from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { useRouter } from "next/navigation";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { Label } from "../ui/label";
import { EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { BattingFilters } from "./batting-filters.component";
import { BowlingFilters } from "./bowling-filters.component";
import { Input } from "../ui/input";
import { PlusSquareIcon, Trash2 } from "lucide-react";

export function StatsFiltersComponent({
  data,
  defaultPlayingFormat,
  defaultStatsType,
  defaultIsMale,
  defaultView,
}: {
  data: IStatsFilters;
  defaultPlayingFormat: string;
  defaultStatsType: EnumStatsType;
  defaultIsMale: "true" | "false";
  defaultView: EnumStatsView;
}) {
  const [statsTypeValue, setStatsTypeValue] = useState(defaultStatsType || "batting");
  const router = useRouter();

  function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const keys = formData.keys().toArray();
    const filters = [`playing_format=${defaultPlayingFormat}`];
    const done = new Set<string>();
    for (const key of keys) {
      if (done.has(key)) continue;
      const values = formData.getAll(key);
      for (const value of values) {
        if (value) {
          filters.push(`${key}=${value}`);
        }
      }
      done.add(key);
    }

    const url = `/stats?${filters.join("&")}`;
    router.push(url);
  }

  return (
    <div className="flex flex-col gap-4">
      <FormatOptions playingFormat={defaultPlayingFormat} />
      <form onSubmit={(e) => handleSubmit(e)} className="px-2 flex flex-col gap-4">
        <MainDropdowns
          defaultIsMale={defaultIsMale}
          defaultStatsType={statsTypeValue}
          setStatsTypeValue={setStatsTypeValue}
        />
        <CommonFilters data={data} />
        {statsTypeValue === "batting" ? (
          <BattingFilters />
        ) : statsTypeValue === "bowling" ? (
          <BowlingFilters />
        ) : (
          <></>
        )}
        <ViewGroupOptions statsType={statsTypeValue} defaultView={defaultView} />
        <PaginationInputs defaultPage={1} defaultLimit={50} />
        <Button type="submit" className="w-24">
          Submit
        </Button>
      </form>
    </div>
  );
}

function MainDropdowns({
  defaultIsMale,
  defaultStatsType,
  setStatsTypeValue,
}: {
  defaultIsMale: "true" | "false";
  defaultStatsType: "batting" | "bowling" | "team";
  setStatsTypeValue: Dispatch<SetStateAction<"batting" | "bowling" | "team">>;
}) {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Gender</p>
        <div>
          <Select name="is_male" defaultValue={defaultIsMale || "true"}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Gender" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Gender</SelectLabel>
                <SelectItem value="true">Male</SelectItem>
                <SelectItem value="false">Female</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Stats Type</p>
        <div>
          <Select
            name="type"
            value={defaultStatsType}
            onValueChange={(value: "batting" | "bowling" | "team") => {
              setStatsTypeValue(value);
            }}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Stats Type" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Type</SelectLabel>
                <SelectItem value="batting">Batting</SelectItem>
                <SelectItem value="bowling">Bowling</SelectItem>
                <SelectItem value="team">Team</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>
    </div>
  );
}

function FormatOptions({ playingFormat }: { playingFormat: string }) {
  return (
    <div className="w-full flex flex-row bg-secondary justify-center gap-2">
      <Link
        href={`/stats/filters?playing_format=Test`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "Test" ? { color: "var(--color-sky-500)" } : {}}
      >
        Tests
      </Link>
      <Link
        href={`/stats/filters?playing_format=ODI`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "ODI" ? { color: "var(--color-sky-500)" } : {}}
      >
        ODIs
      </Link>
      <Link
        href={`/stats/filters?playing_format=T20I`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "T20I" ? { color: "var(--color-sky-500)" } : {}}
      >
        T20Is
      </Link>
      <Link
        href={`/stats/filters?playing_format=first_class`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "first_class" ? { color: "var(--color-sky-500)" } : {}}
      >
        FC
      </Link>
      <Link
        href={`/stats/filters?playing_format=list_a`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "list_a" ? { color: "var(--color-sky-500)" } : {}}
      >
        List A
      </Link>
      <Link
        href={`/stats/filters?playing_format=T20`}
        className="bg-secondary px-2 py-1 rounded"
        style={playingFormat === "T20" ? { color: "var(--color-sky-500)" } : {}}
      >
        T20
      </Link>
    </div>
  );
}

function ViewGroupOptions({
  statsType,
  defaultView,
}: {
  statsType: "batting" | "bowling" | "team";
  defaultView: "overall" | "individual";
}) {
  const commonGroupOptions = {
    overall: [
      { value: "team-innings", label: "Team Innings" },
      { value: "matches", label: "Matches" },
      { value: "teams", label: "Teams" },
      { value: "oppositions", label: "Oppositions" },
      { value: "grounds", label: "Grounds" },
      { value: "host-nations", label: "Host Nations" },
      { value: "continents", label: "Continents" },
      { value: "series", label: "Series" },
      { value: "tournaments", label: "Tournaments" },
      { value: "years", label: "Years" },
      { value: "seasons", label: "Seasons" },
      { value: "decades", label: "Decades" },
      { value: "aggregate", label: "Aggregate" },
    ],
    individual: [
      { value: "innings", label: "Innings" },
      { value: "match-totals", label: "Match Totals" },
      { value: "series", label: "Series" },
      { value: "tournaments", label: "Tournaments" },
      { value: "grounds", label: "Grounds" },
      { value: "host-nations", label: "Host Nations" },
      { value: "oppositions", label: "Oppositions" },
      { value: "years", label: "Years" },
      { value: "seasons", label: "Seasons" },
    ],
  };

  const particularGroupOptions = {
    batting: {
      overall: [{ value: "batters", label: "Batters" }, ...commonGroupOptions["overall"]],
      individual: [...commonGroupOptions["individual"]],
    },
    bowling: {
      overall: [{ value: "bowlers", label: "Bowlers" }, ...commonGroupOptions["overall"]],
      individual: [...commonGroupOptions["individual"]],
    },
    team: {
      overall: [],
      individual: [],
    },
  };

  const qualificationFilters: Record<
    EnumStatsType,
    Record<
      string,
      {
        label: string;
        commonAll: boolean;
        commonOverall: boolean;
        commonIndividual: boolean;
        overallGroups: string[];
        individualGroups: string[];
      }
    >
  > = {
    batting: {
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_batted: { label: "Innings Batted", ...getQualificationFilterObject(true) },
      not_outs: { label: "Not Outs", ...getQualificationFilterObject(true) },
      runs_scored: { label: "Runs Scored", ...getQualificationFilterObject(true) },
      balls_faced: { label: "Balls Faced", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      centuries: { label: "Centuries", ...getQualificationFilterObject(true) },
      half_centuries: { label: "Half Centuries", ...getQualificationFilterObject(true) },
      fifty_plus_scores: { label: "50+ scores", ...getQualificationFilterObject(true) },
      ducks: { label: "Ducks", ...getQualificationFilterObject(true) },
      fours_scored: { label: "Fours scored", ...getQualificationFilterObject(true) },
      sixes_scored: { label: "Sixes scored", ...getQualificationFilterObject(true) },
    },
    bowling: {
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_bowled: { label: "Innings Bowled", ...getQualificationFilterObject(true) },
      overs_bowled: { label: "Overs Bowled", ...getQualificationFilterObject(true) },
      maiden_overs: { label: "Maiden Overs", ...getQualificationFilterObject(true) },
      runs_conceded: { label: "Runs Conceded", ...getQualificationFilterObject(true) },
      wickets_taken: { label: "Wickets Taken", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      economy: { label: "Economy", ...getQualificationFilterObject(true) },
      fours_conceded: { label: "Fours Conceded", ...getQualificationFilterObject(true) },
      sixes_conceded: { label: "Sixes Conceded", ...getQualificationFilterObject(true) },
      four_wkt_hauls: {
        label: "4-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      five_wkt_hauls: {
        label: "5-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      ten_wkt_hauls: {
        label: "10-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
    },
    team: {},
  };

  const sortOptionsMap: Record<
    EnumStatsType,
    Record<
      string,
      {
        label: string;
        commonAll: boolean;
        commonOverall: boolean;
        commonIndividual: boolean;
        overallGroups: string[];
        individualGroups: string[];
      }
    >
  > = {
    batting: {
      runs_scored: { label: "Runs Scored", ...getQualificationFilterObject(true) },
      start_date: { label: "Start Date", ...getQualificationFilterObject(true) },
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_batted: { label: "Innings Batted", ...getQualificationFilterObject(true) },
      not_outs: { label: "Not Outs", ...getQualificationFilterObject(true) },
      balls_faced: { label: "Balls Faced", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      centuries: { label: "Centuries", ...getQualificationFilterObject(true) },
      half_centuries: { label: "Half Centuries", ...getQualificationFilterObject(true) },
      fifty_plus_scores: { label: "50+ scores", ...getQualificationFilterObject(true) },
      ducks: { label: "Ducks", ...getQualificationFilterObject(true) },
      fours_scored: { label: "Fours scored", ...getQualificationFilterObject(true) },
      sixes_scored: { label: "Sixes scored", ...getQualificationFilterObject(true) },
      player_name: { label: "Player Name", ...getQualificationFilterObject(false, false, true, ["batters"]) },
      innings_number: {
        label: "Innings Number",
        ...getQualificationFilterObject(false, false, false, [], ["innings"]),
      },
    },
    bowling: {
      wickets_taken: { label: "Wickets Taken", ...getQualificationFilterObject(true) },
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_bowled: { label: "Innings Bowled", ...getQualificationFilterObject(true) },
      overs_bowled: { label: "Overs Bowled", ...getQualificationFilterObject(true) },
      maiden_overs: { label: "Maiden Overs", ...getQualificationFilterObject(true) },
      runs_conceded: { label: "Runs Conceded", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      economy: { label: "Economy", ...getQualificationFilterObject(true) },
      four_wkt_hauls: {
        label: "4-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      five_wkt_hauls: {
        label: "5-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      ten_wkt_hauls: {
        label: "10-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      best_bowling_match: {
        label: "Best Bowling Match",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      best_bowling_innings: {
        label: "Best Bowling Innings",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      player_name: { label: "Player Name", ...getQualificationFilterObject(false, false, true, ["bowlers"]) },
    },
    team: {},
  };

  const [viewValue, setViewValue] = useState(defaultView || "overall");
  const [groupOptions, setGroupOptions] = useState([...particularGroupOptions[statsType][viewValue]]);
  const [groupValue, setGroupValue] = useState(groupOptions?.[0]?.value);

  const initialQualificationFilters: { label: string; name: string }[] = [];
  for (const filter in qualificationFilters[statsType]) {
    const info = qualificationFilters[statsType][filter];
    if (
      info.commonAll ||
      (viewValue === "overall" && info.commonOverall) ||
      (viewValue === "individual" && info.commonIndividual) ||
      (viewValue === "overall" && info.overallGroups.includes(groupValue)) ||
      (viewValue === "individual" && info.individualGroups.includes(groupValue))
    ) {
      initialQualificationFilters.push({ name: filter, label: info.label });
    }
  }

  const initialSortOptions: { label: string; name: string }[] = [];
  for (const filter in sortOptionsMap[statsType]) {
    const info = sortOptionsMap[statsType][filter];
    if (
      info.commonAll ||
      (viewValue === "overall" && info.commonOverall) ||
      (viewValue === "individual" && info.commonIndividual) ||
      (viewValue === "overall" && info.overallGroups.includes(groupValue)) ||
      (viewValue === "individual" && info.individualGroups.includes(groupValue))
    ) {
      initialSortOptions.push({ name: filter, label: info.label });
    }
  }

  const [qualificationFilterOptions, setQualificationFilterOptions] = useState(initialQualificationFilters);
  const [usedQualificationFilterOptions, setUsedQualificationFilterOptions] = useState<
    { name: string; label: string }[]
  >([qualificationFilterOptions[0]]);
  const [sortOptions, setSortOptions] = useState(initialSortOptions);
  const [sortByValue, setSortByValue] = useState(sortOptions?.[0].name);

  useEffect(() => {
    const options = particularGroupOptions[statsType][viewValue];
    setGroupOptions([...options]);
  }, [statsType, viewValue]);

  useEffect(() => {
    setGroupValue(groupOptions?.[0]?.value);
  }, [groupOptions]);

  useEffect(() => {
    const arr: { label: string; name: string }[] = [];
    for (const filter in qualificationFilters[statsType]) {
      const info = qualificationFilters[statsType][filter];
      if (
        info.commonAll ||
        (viewValue === "overall" && info.commonOverall) ||
        (viewValue === "individual" && info.commonIndividual) ||
        (viewValue === "overall" && info.overallGroups.includes(groupValue)) ||
        (viewValue === "individual" && info.individualGroups.includes(groupValue))
      ) {
        arr.push({ name: filter, label: info.label });
      }
    }
    setQualificationFilterOptions(arr);

    const arr2: { label: string; name: string }[] = [];
    for (const filter in sortOptionsMap[statsType]) {
      const info = sortOptionsMap[statsType][filter];
      if (
        info.commonAll ||
        (viewValue === "overall" && info.commonOverall) ||
        (viewValue === "individual" && info.commonIndividual) ||
        (viewValue === "overall" && info.overallGroups.includes(groupValue)) ||
        (viewValue === "individual" && info.individualGroups.includes(groupValue))
      ) {
        arr2.push({ name: filter, label: info.label });
      }
    }
    setSortOptions(arr2);
  }, [groupValue]);

  useEffect(() => {
    setUsedQualificationFilterOptions([qualificationFilterOptions[0]]);
  }, [qualificationFilterOptions]);

  useEffect(() => {
    setSortByValue(sortOptions[0]?.name);
  }, [sortOptions]);

  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>View Format</p>
        <RadioGroup
          name="view"
          value={viewValue}
          className="flex gap-4"
          onValueChange={(value: "overall" | "individual") => {
            setViewValue(value);
          }}
        >
          <div className="flex items-center gap-1">
            <RadioGroupItem value="overall" id="overall" />
            <Label htmlFor="overall" className="capitalize font-normal">
              overall
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="individual" id="individual" />
            <Label htmlFor="individual" className="capitalize font-normal">
              individual
            </Label>
          </div>
        </RadioGroup>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Group By</p>
        <Select
          name="group"
          value={groupValue}
          onValueChange={(value: string) => {
            setGroupValue(value);
          }}
        >
          <SelectTrigger className="w-[180px]">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel className="text-sm">Group</SelectLabel>
              {groupOptions.map((item) => (
                <SelectItem
                  key={item.value}
                  value={item.value}
                  onSelect={() => {
                    setGroupValue(item.value);
                  }}
                >
                  {item.label}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <div className="hidden md:flex justify-between">
        <div className="flex gap-4">
          <p style={{ minWidth: "175px" }}>Result Qualifications</p>
          <div className="flex flex-col gap-2">
            {usedQualificationFilterOptions?.length ? (
              usedQualificationFilterOptions?.map((item, idx) => {
                return (
                  <div key={item.name} className="flex gap-2">
                    <Select
                      defaultValue={item.name}
                      onValueChange={(value) => {
                        const updatedUsedOptions = [...usedQualificationFilterOptions];
                        const newOption = qualificationFilterOptions.find((x) => x.name === value);
                        if (newOption) {
                          updatedUsedOptions[idx] = { ...newOption };
                        }
                        setUsedQualificationFilterOptions(updatedUsedOptions);
                      }}
                    >
                      <SelectTrigger className="w-[180px]">
                        <SelectValue placeholder="Select" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectGroup>
                          <SelectLabel className="text-sm">Qualification Filters</SelectLabel>
                          {qualificationFilterOptions.map((x) => {
                            const isUsed = usedQualificationFilterOptions.find((y) => y.name === x.name);
                            if (!isUsed || x.name === item.name) {
                              return (
                                <SelectItem key={x.name} value={x.name}>
                                  {x.label}
                                </SelectItem>
                              );
                            }
                          })}
                        </SelectGroup>
                      </SelectContent>
                    </Select>

                    <Input
                      type="number"
                      name={`min__${item.name}`}
                      placeholder="Min"
                      min={1}
                      step={1}
                      className="w-24"
                    />
                    <Input
                      type="number"
                      name={`max__${item.name}`}
                      placeholder="Max"
                      min={1}
                      step={1}
                      className="w-24"
                    />
                    <Button
                      type="button"
                      onClick={() => {
                        const updatedFilters = usedQualificationFilterOptions.filter(
                          (x) => x.name !== item.name
                        );
                        setUsedQualificationFilterOptions(updatedFilters);
                      }}
                    >
                      <Trash2 />
                    </Button>
                  </div>
                );
              })
            ) : (
              <p>Add a qualification Filter</p>
            )}
          </div>
        </div>
        <div>
          <Button
            type="button"
            className="flex gap-2"
            disabled={qualificationFilterOptions?.length === usedQualificationFilterOptions?.length}
            onClick={() => {
              const arr = [...usedQualificationFilterOptions];
              for (const option of qualificationFilterOptions) {
                const isUsed = usedQualificationFilterOptions.find((item) => item.name === option.name);
                if (!isUsed) {
                  arr.push({ ...option });
                  break;
                }
              }
              setUsedQualificationFilterOptions(arr);
            }}
          >
            <PlusSquareIcon /> Add
          </Button>
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Sort By</p>
        <Select
          name="sort_by"
          value={sortByValue}
          onValueChange={(value) => {
            setSortByValue(value);
          }}
        >
          <SelectTrigger className="w-[180px]">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel className="text-sm">Sort By</SelectLabel>
              {sortOptions.map((item) => (
                <SelectItem key={item.name} value={item.name}>
                  {item.label}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Sord Order</p>
        <RadioGroup name="sort_order" defaultValue="default" className="flex gap-4">
          <div className="flex items-center gap-1">
            <RadioGroupItem value="default" id="sort-default" />
            <Label htmlFor="sort-default" className="capitalize font-normal">
              default
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="reverse" id="sort-reverse" />
            <Label htmlFor="sort-reverse" className="capitalize font-normal">
              reverse
            </Label>
          </div>
        </RadioGroup>
      </div>
    </div>
  );
}

function PaginationInputs({ defaultPage, defaultLimit }: { defaultPage: number; defaultLimit: number }) {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Page No.</p>
        <Input
          type="number"
          name="__page"
          min={1}
          step={1}
          className="w-24"
          defaultValue={defaultPage || 1}
        />
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Records per Page</p>
        <div>
          <Select name="__limit" defaultValue={defaultLimit.toString() || "50"}>
            <SelectTrigger className="w-24">
              <SelectValue placeholder="50" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Records per Page</SelectLabel>
                <SelectItem value="50">50</SelectItem>
                <SelectItem value="100">100</SelectItem>
                <SelectItem value="150">150</SelectItem>
                <SelectItem value="200">200</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>
    </div>
  );
}

function getQualificationFilterObject(
  commonAll = false,
  commonOverall = false,
  commonIndividual = false,
  overallGroups: string[] = [],
  individualGroups: string[] = []
) {
  const obj = {
    commonAll,
    commonOverall,
    commonIndividual,
    overallGroups,
    individualGroups,
  };

  if (commonAll) {
    obj.commonOverall = obj.commonIndividual = false;
    obj.overallGroups = [];
    obj.individualGroups = [];
  }

  if (obj.commonOverall) obj.overallGroups = [];
  if (obj.commonIndividual) obj.individualGroups = [];

  return obj;
}
