--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Ubuntu 16.4-1.pgdg22.04+2)
-- Dumped by pg_dump version 17.0 (Ubuntu 17.0-1.pgdg22.04+1)

-- Started on 2024-10-22 20:44:55 IST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 852 (class 1247 OID 20575)
-- Name: career_stats; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.career_stats AS (
	matches_played integer,
	innings_batted integer,
	runs_scored integer,
	batting_dismissals integer,
	balls_faced integer,
	fours_scored integer,
	sixes_scored integer,
	centuries_scored integer,
	fifties_scored integer,
	highest_score integer,
	is_highest_not_out boolean,
	innings_bowled integer,
	runs_conceded integer,
	wickets_taken integer,
	balls_bowled integer,
	fours_conceded integer,
	sixes_conceded integer,
	four_wkt_hauls integer,
	five_wkt_hauls integer,
	ten_wkt_hauls integer,
	best_inn_fig_runs integer,
	best_inn_fig_wkts integer,
	best_match_fig_runs integer,
	best_match_fig_wkts integer,
	debut_match_id integer,
	last_match_id integer
);


ALTER TYPE public.career_stats OWNER TO postgres;

--
-- TOC entry 224 (class 1255 OID 20607)
-- Name: get_player_profile_by_id(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_player_profile_by_id(player_id integer) RETURNS TABLE(id integer, name text, playing_role text, nationality text, is_male boolean, date_of_birth date, image_url text, biography text, batting_styles text[], primary_batting_style text, bowling_styles text[], primary_bowling_style text, teams_represented jsonb, test_stats jsonb, odi_stats jsonb, t20i_stats jsonb, fc_stats jsonb, lista_stats jsonb, t20_stats jsonb, cricsheet_id text, cricinfo_id text, cricbuzz_id text)
    LANGUAGE plpgsql
    AS $$ BEGIN RETURN QUERY
SELECT p.id,
    p.name,
    p.playing_role,
    p.nationality,
    p.is_male,
    p.date_of_birth,
    p.image_url,
    p.biography,
    p.batting_styles,
    p.primary_batting_style,
    p.bowling_styles,
    p.primary_bowling_style,
    COALESCE(
        jsonb_agg(jsonb_build_object('id', t.id, 'name', t.name)) FILTER (
            WHERE t.id IS NOT NULL
        ),
        '[]'
    ) AS teams_represented,
    jsonb_build_object(
        'matches_played',
        (p.test_stats).matches_played,
        'innings_batted',
        (p.test_stats).innings_batted,
        'runs_scored',
        (p.test_stats).runs_scored,
        'batting_dismissals',
        (p.test_stats).batting_dismissals,
        'balls_faced',
        (p.test_stats).balls_faced,
        'fours_scored',
        (p.test_stats).fours_scored,
        'sixes_scored',
        (p.test_stats).sixes_scored,
        'centuries_scored',
        (p.test_stats).centuries_scored,
        'fifties_scored',
        (p.test_stats).fifties_scored,
        'highest_score',
        (p.test_stats).highest_score,
        'is_highest_not_out',
        (p.test_stats).is_highest_not_out,
        'innings_bowled',
        (p.test_stats).innings_bowled,
        'wickets_taken',
        (p.test_stats).wickets_taken,
        'runs_conceded',
        (p.test_stats).runs_conceded,
        'balls_bowled',
        (p.test_stats).balls_bowled,
        'fours_conceded',
        (p.test_stats).fours_conceded,
        'sixes_conceded',
        (p.test_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.test_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.test_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.test_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.test_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.test_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.test_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.test_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            test_dm.id,
            'team1',
            test_dm_team1.name,
            'team2',
            test_dm_team2.name,
            'date',
            test_dm.start_datetime,
            'ground',
            test_dm_ground.name
        ),
        'last_match',
        jsonb_build_object(
            'id',
            test_lm.id,
            'team1',
            test_lm_team1.id,
            'team2',
            test_lm_team2.id,
            'date',
            test_lm.start_datetime,
            'ground',
            test_lm.ground_id
        )
    ) AS test_stats,
    jsonb_build_object(
        'matches_played',
        (p.odi_stats).matches_played,
        'innings_batted',
        (p.odi_stats).innings_batted,
        'runs_scored',
        (p.odi_stats).runs_scored,
        'batting_dismissals',
        (p.odi_stats).batting_dismissals,
        'balls_faced',
        (p.odi_stats).balls_faced,
        'fours_scored',
        (p.odi_stats).fours_scored,
        'sixes_scored',
        (p.odi_stats).sixes_scored,
        'centuries_scored',
        (p.odi_stats).centuries_scored,
        'fifties_scored',
        (p.odi_stats).fifties_scored,
        'highest_score',
        (p.odi_stats).highest_score,
        'is_highest_not_out',
        (p.odi_stats).is_highest_not_out,
        'innings_bowled',
        (p.odi_stats).innings_bowled,
        'wickets_taken',
        (p.odi_stats).wickets_taken,
        'runs_conceded',
        (p.odi_stats).runs_conceded,
        'balls_bowled',
        (p.odi_stats).balls_bowled,
        'fours_conceded',
        (p.odi_stats).fours_conceded,
        'sixes_conceded',
        (p.odi_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.odi_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.odi_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.odi_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.odi_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.odi_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.odi_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.odi_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            odi_dm.id,
            'team1',
            odi_dm_team1.name,
            'team2',
            odi_dm_team2.name,
            'date',
            odi_dm.start_datetime,
            'ground',
            odi_dm.ground_id
        ),
        'last_match',
        jsonb_build_object(
            'id',
            odi_lm.id,
            'team1',
            odi_lm_team1.id,
            'team2',
            odi_lm_team2.id,
            'date',
            odi_lm.start_datetime,
            'ground',
            odi_lm.ground_id
        )
    ) AS odi_stats,
    jsonb_build_object(
        'matches_played',
        (p.t20i_stats).matches_played,
        'innings_batted',
        (p.t20i_stats).innings_batted,
        'runs_scored',
        (p.t20i_stats).runs_scored,
        'batting_dismissals',
        (p.t20i_stats).batting_dismissals,
        'balls_faced',
        (p.t20i_stats).balls_faced,
        'fours_scored',
        (p.t20i_stats).fours_scored,
        'sixes_scored',
        (p.t20i_stats).sixes_scored,
        'centuries_scored',
        (p.t20i_stats).centuries_scored,
        'fifties_scored',
        (p.t20i_stats).fifties_scored,
        'highest_score',
        (p.t20i_stats).highest_score,
        'is_highest_not_out',
        (p.t20i_stats).is_highest_not_out,
        'innings_bowled',
        (p.t20i_stats).innings_bowled,
        'wickets_taken',
        (p.t20i_stats).wickets_taken,
        'runs_conceded',
        (p.t20i_stats).runs_conceded,
        'balls_bowled',
        (p.t20i_stats).balls_bowled,
        'fours_conceded',
        (p.t20i_stats).fours_conceded,
        'sixes_conceded',
        (p.t20i_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.t20i_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.t20i_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.t20i_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.t20i_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.t20i_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.t20i_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.t20i_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            t20i_dm.id,
            'team1',
            t20i_dm_team1.name,
            'team2',
            t20i_dm_team2.name,
            'date',
            t20i_dm.start_datetime,
            'ground',
            t20i_dm.ground_id
        ),
        'last_match',
        jsonb_build_object(
            'id',
            t20i_lm.id,
            'team1',
            t20i_lm_team1.id,
            'team2',
            t20i_lm_team2.id,
            'date',
            t20i_lm.start_datetime,
            'ground',
            t20i_lm.ground_id
        )
    ) AS t20i_stats,
    jsonb_build_object(
        'matches_played',
        (p.fc_stats).matches_played,
        'innings_batted',
        (p.fc_stats).innings_batted,
        'runs_scored',
        (p.fc_stats).runs_scored,
        'batting_dismissals',
        (p.fc_stats).batting_dismissals,
        'balls_faced',
        (p.fc_stats).balls_faced,
        'fours_scored',
        (p.fc_stats).fours_scored,
        'sixes_scored',
        (p.fc_stats).sixes_scored,
        'centuries_scored',
        (p.fc_stats).centuries_scored,
        'fifties_scored',
        (p.fc_stats).fifties_scored,
        'highest_score',
        (p.fc_stats).highest_score,
        'is_highest_not_out',
        (p.fc_stats).is_highest_not_out,
        'innings_bowled',
        (p.fc_stats).innings_bowled,
        'wickets_taken',
        (p.fc_stats).wickets_taken,
        'runs_conceded',
        (p.fc_stats).runs_conceded,
        'balls_bowled',
        (p.fc_stats).balls_bowled,
        'fours_conceded',
        (p.fc_stats).fours_conceded,
        'sixes_conceded',
        (p.fc_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.fc_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.fc_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.fc_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.fc_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.fc_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.fc_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.fc_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            fc_dm.id,
            'team1',
            fc_dm_team1.name,
            'team2',
            fc_dm_team2.name,
            'date',
            fc_dm.start_datetime,
            'ground',
            fc_dm.ground_id
        ),
        'last_match',
        jsonb_build_object(
            'id',
            fc_lm.id,
            'team1',
            fc_lm_team1.id,
            'team2',
            fc_lm_team2.id,
            'date',
            fc_lm.start_datetime,
            'ground',
            fc_lm.ground_id
        )
    ) AS fc_stats,
    jsonb_build_object(
        'matches_played',
        (p.lista_stats).matches_played,
        'innings_batted',
        (p.lista_stats).innings_batted,
        'runs_scored',
        (p.lista_stats).runs_scored,
        'batting_dismissals',
        (p.lista_stats).batting_dismissals,
        'balls_faced',
        (p.lista_stats).balls_faced,
        'fours_scored',
        (p.lista_stats).fours_scored,
        'sixes_scored',
        (p.lista_stats).sixes_scored,
        'centuries_scored',
        (p.lista_stats).centuries_scored,
        'fifties_scored',
        (p.lista_stats).fifties_scored,
        'highest_score',
        (p.lista_stats).highest_score,
        'is_highest_not_out',
        (p.lista_stats).is_highest_not_out,
        'innings_bowled',
        (p.lista_stats).innings_bowled,
        'wickets_taken',
        (p.lista_stats).wickets_taken,
        'runs_conceded',
        (p.lista_stats).runs_conceded,
        'balls_bowled',
        (p.lista_stats).balls_bowled,
        'fours_conceded',
        (p.lista_stats).fours_conceded,
        'sixes_conceded',
        (p.lista_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.lista_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.lista_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.lista_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.lista_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.lista_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.lista_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.lista_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            lista_dm.id,
            'team1',
            lista_dm_team1.name,
            'team2',
            lista_dm_team2.name,
            'date',
            lista_dm.start_datetime,
            'ground',
            lista_dm.ground_id
        ),
        'last_match',
        jsonb_build_object(
            'id',
            lista_lm.id,
            'team1',
            lista_lm_team1.id,
            'team2',
            lista_lm_team2.id,
            'date',
            lista_lm.start_datetime,
            'ground',
            lista_lm.ground_id
        )
    ) AS lista_stats,
    jsonb_build_object(
        'matches_played',
        (p.t20_stats).matches_played,
        'innings_batted',
        (p.t20_stats).innings_batted,
        'runs_scored',
        (p.t20_stats).runs_scored,
        'batting_dismissals',
        (p.t20_stats).batting_dismissals,
        'balls_faced',
        (p.t20_stats).balls_faced,
        'fours_scored',
        (p.t20_stats).fours_scored,
        'sixes_scored',
        (p.t20_stats).sixes_scored,
        'centuries_scored',
        (p.t20_stats).centuries_scored,
        'fifties_scored',
        (p.t20_stats).fifties_scored,
        'highest_score',
        (p.t20_stats).highest_score,
        'is_highest_not_out',
        (p.t20_stats).is_highest_not_out,
        'innings_bowled',
        (p.t20_stats).innings_bowled,
        'wickets_taken',
        (p.t20_stats).wickets_taken,
        'runs_conceded',
        (p.t20_stats).runs_conceded,
        'balls_bowled',
        (p.t20_stats).balls_bowled,
        'fours_conceded',
        (p.t20_stats).fours_conceded,
        'sixes_conceded',
        (p.t20_stats).sixes_conceded,
        'four_wkt_hauls',
        (p.t20_stats).four_wkt_hauls,
        'five_wkt_hauls',
        (p.t20_stats).five_wkt_hauls,
        'ten_wkt_hauls',
        (p.t20_stats).ten_wkt_hauls,
        'best_inn_fig_runs',
        (p.t20_stats).best_inn_fig_runs,
        'best_inn_fig_wkts',
        (p.t20_stats).best_inn_fig_wkts,
        'best_match_fig_runs',
        (p.t20_stats).best_match_fig_runs,
        'best_match_fig_wkts',
        (p.t20_stats).best_match_fig_wkts,
        'debut_match',
        jsonb_build_object(
            'id',
            t20_dm.id,
            'team1',
            t20_dm_team1.name,
            'team2',
            t20_dm_team2.name,
            'date',
            t20_dm.start_datetime,
            'ground',
            t20_dm.ground_id
        ),
        'last_match',
        jsonb_build_object(
            'id',
            t20_lm.id,
            'team1',
            t20_lm_team1.id,
            'team2',
            t20_lm_team2.id,
            'date',
            t20_lm.start_datetime,
            'ground',
            t20_lm.ground_id
        )
    ) AS t20_stats,
    p.cricsheet_id,
    p.cricinfo_id,
    p.cricbuzz_id
FROM players p
    LEFT JOIN LATERAL unnest(p.teams_represented_id) AS team_id ON true
    LEFT JOIN teams t ON t.id = team_id
    LEFT JOIN matches test_dm ON test_dm.id = (p.test_stats).debut_match_id
    LEFT JOIN matches test_lm ON test_lm.id = (p.test_stats).last_match_id
    LEFT JOIN matches odi_dm ON odi_dm.id = (p.odi_stats).debut_match_id
    LEFT JOIN matches odi_lm ON odi_lm.id = (p.odi_stats).last_match_id
    LEFT JOIN matches t20i_dm ON t20i_dm.id = (p.t20i_stats).debut_match_id
    LEFT JOIN matches t20i_lm ON t20i_lm.id = (p.t20i_stats).last_match_id
    LEFT JOIN matches fc_dm ON fc_dm.id = (p.fc_stats).debut_match_id
    LEFT JOIN matches fc_lm ON fc_lm.id = (p.fc_stats).last_match_id
    LEFT JOIN matches lista_dm ON lista_dm.id = (p.lista_stats).debut_match_id
    LEFT JOIN matches lista_lm ON lista_lm.id = (p.lista_stats).last_match_id
    LEFT JOIN matches t20_dm ON t20_dm.id = (p.t20_stats).debut_match_id
    LEFT JOIN matches t20_lm ON t20_lm.id = (p.t20_stats).last_match_id
    LEFT JOIN teams test_dm_team1 ON test_dm.team1_id = test_dm_team1.id
    LEFT JOIN teams test_lm_team1 ON test_lm.team1_id = test_lm_team1.id
    LEFT JOIN teams test_dm_team2 ON test_dm.team2_id = test_dm_team2.id
    LEFT JOIN teams test_lm_team2 ON test_lm.team2_id = test_lm_team2.id
    LEFT JOIN teams odi_dm_team1 ON odi_dm.team1_id = odi_dm_team1.id
    LEFT JOIN teams odi_lm_team1 ON odi_lm.team1_id = odi_lm_team1.id
    LEFT JOIN teams odi_dm_team2 ON odi_dm.team2_id = odi_dm_team2.id
    LEFT JOIN teams odi_lm_team2 ON odi_lm.team2_id = odi_lm_team2.id
    LEFT JOIN teams t20i_dm_team1 ON t20i_dm.team1_id = t20i_dm_team1.id
    LEFT JOIN teams t20i_lm_team1 ON t20i_lm.team1_id = t20i_lm_team1.id
    LEFT JOIN teams t20i_dm_team2 ON t20i_dm.team2_id = t20i_dm_team2.id
    LEFT JOIN teams t20i_lm_team2 ON t20i_lm.team2_id = t20i_lm_team2.id
    LEFT JOIN teams fc_dm_team1 ON fc_dm.team1_id = fc_dm_team1.id
    LEFT JOIN teams fc_lm_team1 ON fc_lm.team1_id = fc_lm_team1.id
    LEFT JOIN teams fc_dm_team2 ON fc_dm.team2_id = fc_dm_team2.id
    LEFT JOIN teams fc_lm_team2 ON fc_lm.team2_id = fc_lm_team2.id
    LEFT JOIN teams lista_dm_team1 ON lista_dm.team1_id = lista_dm_team1.id
    LEFT JOIN teams lista_lm_team1 ON lista_lm.team1_id = lista_lm_team1.id
    LEFT JOIN teams lista_dm_team2 ON lista_dm.team2_id = lista_dm_team2.id
    LEFT JOIN teams lista_lm_team2 ON lista_lm.team2_id = lista_lm_team2.id
    LEFT JOIN teams t20_dm_team1 ON t20_dm.team1_id = t20_dm_team1.id
    LEFT JOIN teams t20_lm_team1 ON t20_lm.team1_id = t20_lm_team1.id
    LEFT JOIN teams t20_dm_team2 ON t20_dm.team2_id = t20_dm_team2.id
    LEFT JOIN teams t20_lm_team2 ON t20_lm.team2_id = t20_lm_team2.id
    LEFT JOIN grounds test_dm_ground ON test_dm.ground_id = test_dm_ground.id
    LEFT JOIN grounds odi_dm_ground ON odi_dm.ground_id = odi_dm_ground.id
    LEFT JOIN grounds t20i_dm_ground ON t20i_dm.ground_id = t20i_dm_ground.id
    LEFT JOIN grounds fc_dm_ground ON fc_dm.ground_id = fc_dm_ground.id
    LEFT JOIN grounds lista_dm_ground ON lista_dm.ground_id = lista_dm_ground.id
    LEFT JOIN grounds t20_dm_ground ON t20_dm.ground_id = t20_dm_ground.id
    LEFT JOIN grounds test_lm_ground ON test_lm.ground_id = test_lm_ground.id
    LEFT JOIN grounds odi_lm_ground ON odi_lm.ground_id = odi_lm_ground.id
    LEFT JOIN grounds t20i_lm_ground ON t20i_lm.ground_id = t20i_lm_ground.id
    LEFT JOIN grounds fc_lm_ground ON fc_lm.ground_id = fc_lm_ground.id
    LEFT JOIN grounds lista_lm_ground ON lista_lm.ground_id = lista_lm_ground.id
    LEFT JOIN grounds t20_lm_ground ON t20_lm.ground_id = t20_lm_ground.id
WHERE p.id = player_id
GROUP BY p.id,
    test_dm.id,
    test_dm_team1.id,
    test_dm_team1.name,
    test_dm_team2.id,
    test_dm_team2.name,
    test_dm_ground.id,
    test_dm_ground.name,
    test_lm.id,
    test_lm_team1.id,
    test_lm_team1.name,
    test_lm_team2.id,
    test_lm_team2.name,
    test_lm_ground.id,
    test_lm_ground.name,
    odi_dm.id,
    odi_dm_team1.id,
    odi_dm_team1.name,
    odi_dm_team2.id,
    odi_dm_team2.name,
    odi_dm_ground.id,
    odi_dm_ground.name,
    odi_lm.id,
    odi_lm_team1.id,
    odi_lm_team1.name,
    odi_lm_team2.id,
    odi_lm_team2.name,
    odi_lm_ground.id,
    odi_lm_ground.name,
    t20i_dm.id,
    t20i_dm_team1.id,
    t20i_dm_team1.name,
    t20i_dm_team2.id,
    t20i_dm_team2.name,
    t20i_dm_ground.id,
    t20i_dm_ground.name,
    t20i_lm.id,
    t20i_lm_team1.id,
    t20i_lm_team1.name,
    t20i_lm_team2.id,
    t20i_lm_team2.name,
    t20i_lm_ground.id,
    t20i_lm_ground.name,
    fc_dm.id,
    fc_dm_team1.id,
    fc_dm_team1.name,
    fc_dm_team2.id,
    fc_dm_team2.name,
    fc_dm_ground.id,
    fc_dm_ground.name,
    fc_lm.id,
    fc_lm_team1.id,
    fc_lm_team1.name,
    fc_lm_team2.id,
    fc_lm_team2.name,
    fc_lm_ground.id,
    fc_lm_ground.name,
    lista_dm.id,
    lista_dm_team1.id,
    lista_dm_team1.name,
    lista_dm_team2.id,
    lista_dm_team2.name,
    lista_dm_ground.id,
    lista_dm_ground.name,
    lista_lm.id,
    lista_lm_team1.id,
    lista_lm_team1.name,
    lista_lm_team2.id,
    lista_lm_team2.name,
    lista_lm_ground.id,
    lista_lm_ground.name,
    t20_dm.id,
    t20_dm_team1.id,
    t20_dm_team1.name,
    t20_dm_team2.id,
    t20_dm_team2.name,
    t20_dm_ground.id,
    t20_dm_ground.name,
    t20_lm.id,
    t20_lm_team1.id,
    t20_lm_team1.name,
    t20_lm_team2.id,
    t20_lm_team2.name,
    t20_dm_ground.id,
    t20_dm_ground.name;
END;
$$;


ALTER FUNCTION public.get_player_profile_by_id(player_id integer) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 223 (class 1259 OID 20599)
-- Name: grounds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.grounds (
    id integer NOT NULL,
    name text NOT NULL,
    host_nation_id integer NOT NULL,
    city_id integer NOT NULL
);


ALTER TABLE public.grounds OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 20598)
-- Name: grounds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.grounds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.grounds_id_seq OWNER TO postgres;

--
-- TOC entry 3410 (class 0 OID 0)
-- Dependencies: 222
-- Name: grounds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.grounds_id_seq OWNED BY public.grounds.id;


--
-- TOC entry 221 (class 1259 OID 20590)
-- Name: matches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    start_datetime timestamp with time zone NOT NULL,
    team1_id integer NOT NULL,
    team2_id integer NOT NULL,
    is_male boolean NOT NULL,
    tournament_id integer,
    series_id integer,
    host_nation_id integer,
    continent_id integer,
    ground_id integer,
    current_status text,
    final_result text,
    home_team_id integer,
    away_team_id integer,
    match_type text,
    playing_level text,
    playing_format text,
    season_id text,
    is_day_night boolean,
    is_dls boolean,
    toss_winner_team_id integer,
    toss_loser_team_id integer,
    toss_decision text,
    match_winner_team_id integer,
    match_loser_team_id integer,
    is_won_by_runs boolean,
    win_margin integer,
    balls_remaining_after_win integer,
    potm_id integer,
    scorers_id integer[],
    commentators_id integer[]
);


ALTER TABLE public.matches OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 20589)
-- Name: matches_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.matches_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.matches_id_seq OWNER TO postgres;

--
-- TOC entry 3411 (class 0 OID 0)
-- Dependencies: 220
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
-- TOC entry 219 (class 1259 OID 20577)
-- Name: players; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.players (
    id integer NOT NULL,
    name text NOT NULL,
    playing_role text,
    nationality text NOT NULL,
    is_male boolean NOT NULL,
    date_of_birth date NOT NULL,
    image_url text,
    biography text,
    batting_styles text[],
    primary_batting_style text,
    bowling_styles text[],
    primary_bowling_style text,
    teams_represented_id integer[],
    test_stats public.career_stats,
    odi_stats public.career_stats,
    t20i_stats public.career_stats,
    fc_stats public.career_stats,
    lista_stats public.career_stats,
    t20_stats public.career_stats,
    cricsheet_id text,
    cricinfo_id text,
    cricbuzz_id text,
    CONSTRAINT valid_birth_date CHECK ((date_of_birth < CURRENT_DATE)),
    CONSTRAINT valid_image_url CHECK (((image_url IS NULL) OR (image_url ~ '^https?://[^\s/$.?#].[^\s]*$'::text))),
    CONSTRAINT valid_name CHECK ((length(TRIM(BOTH FROM name)) > 0))
);


ALTER TABLE public.players OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 20576)
-- Name: players_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.players_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.players_id_seq OWNER TO postgres;

--
-- TOC entry 3412 (class 0 OID 0)
-- Dependencies: 218
-- Name: players_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.players_id_seq OWNED BY public.players.id;


--
-- TOC entry 216 (class 1259 OID 20564)
-- Name: teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.teams (
    id integer NOT NULL,
    name text NOT NULL,
    is_male boolean DEFAULT true,
    image_url text,
    playing_level text
);


ALTER TABLE public.teams OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 20563)
-- Name: teams_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.teams_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.teams_id_seq OWNER TO postgres;

--
-- TOC entry 3413 (class 0 OID 0)
-- Dependencies: 215
-- Name: teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.teams_id_seq OWNED BY public.teams.id;


--
-- TOC entry 3242 (class 2604 OID 20602)
-- Name: grounds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds ALTER COLUMN id SET DEFAULT nextval('public.grounds_id_seq'::regclass);


--
-- TOC entry 3241 (class 2604 OID 20593)
-- Name: matches id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches ALTER COLUMN id SET DEFAULT nextval('public.matches_id_seq'::regclass);


--
-- TOC entry 3240 (class 2604 OID 20580)
-- Name: players id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players ALTER COLUMN id SET DEFAULT nextval('public.players_id_seq'::regclass);


--
-- TOC entry 3238 (class 2604 OID 20567)
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- TOC entry 3404 (class 0 OID 20599)
-- Dependencies: 223
-- Data for Name: grounds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.grounds (id, name, host_nation_id, city_id) FROM stdin;
\.


--
-- TOC entry 3402 (class 0 OID 20590)
-- Dependencies: 221
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.matches (id, start_datetime, team1_id, team2_id, is_male, tournament_id, series_id, host_nation_id, continent_id, ground_id, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season_id, is_day_night, is_dls, toss_winner_team_id, toss_loser_team_id, toss_decision, match_winner_team_id, match_loser_team_id, is_won_by_runs, win_margin, balls_remaining_after_win, potm_id, scorers_id, commentators_id) FROM stdin;
\.


--
-- TOC entry 3400 (class 0 OID 20577)
-- Dependencies: 219
-- Data for Name: players; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.players (id, name, playing_role, nationality, is_male, date_of_birth, image_url, biography, batting_styles, primary_batting_style, bowling_styles, primary_bowling_style, teams_represented_id, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id) FROM stdin;
2	Rohit Sharma	Top-order Batter	India	t	1993-04-30	\N	\N	\N	\N	\N	\N	\N	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	\N	\N	\N
1	Virat Kohli	Top-order Batter	India	t	1994-11-05	\N	\N	\N	\N	\N	\N	{1,3}	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	\N	\N	\N
\.


--
-- TOC entry 3398 (class 0 OID 20564)
-- Dependencies: 216
-- Data for Name: teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.teams (id, name, is_male, image_url, playing_level) FROM stdin;
1	India	t	\N	international
2	Australia	t	\N	international
3	Royal Challengers Bangalore	t	\N	domestic
\.


--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 222
-- Name: grounds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.grounds_id_seq', 1, false);


--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 220
-- Name: matches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.matches_id_seq', 1, false);


--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 218
-- Name: players_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.players_id_seq', 2, true);


--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 215
-- Name: teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.teams_id_seq', 3, true);


--
-- TOC entry 3253 (class 2606 OID 20606)
-- Name: grounds grounds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds
    ADD CONSTRAINT grounds_pkey PRIMARY KEY (id);


--
-- TOC entry 3251 (class 2606 OID 20597)
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- TOC entry 3249 (class 2606 OID 20587)
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);


--
-- TOC entry 3247 (class 2606 OID 20572)
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (id);


-- Completed on 2024-10-22 20:44:55 IST

--
-- PostgreSQL database dump complete
--

