--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Ubuntu 16.4-1.pgdg22.04+2)
-- Dumped by pg_dump version 17.0 (Ubuntu 17.0-1.pgdg22.04+1)

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
-- Name: article_category; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.article_category AS ENUM (
    'news',
    'preview',
    'analysis',
    'feature',
    'interview',
    'report'
);


ALTER TYPE public.article_category OWNER TO postgres;

--
-- Name: article_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.article_status AS ENUM (
    'published',
    'drafted',
    'trashed'
);


ALTER TYPE public.article_status OWNER TO postgres;

--
-- Name: bowling_style; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.bowling_style AS ENUM (
    'right_arm_fast',
    'left_arm_fast',
    'right_arm_off_break',
    'left_arm_orthodox',
    'right_arm_leg_spin',
    'left_arm_wrist_spin'
);


ALTER TYPE public.bowling_style OWNER TO postgres;

--
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
-- Name: dismissal_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.dismissal_type AS ENUM (
    'caught',
    'bowled',
    'lbw',
    'run_out',
    'stumped',
    'hit_wicket',
    'handed_the_ball',
    'obstructing_the_field',
    'timed_out',
    'retired_hurt',
    'hit_the_ball_twice'
);


ALTER TYPE public.dismissal_type OWNER TO postgres;

--
-- Name: innings_end; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.innings_end AS ENUM (
    'all_out',
    'declared',
    'target_reached',
    'forfeited',
    'incomplete'
);


ALTER TYPE public.innings_end OWNER TO postgres;

--
-- Name: match_final_result; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.match_final_result AS ENUM (
    'winner_decided',
    'tied',
    'drawn',
    'no_result',
    'abandoned'
);


ALTER TYPE public.match_final_result OWNER TO postgres;

--
-- Name: match_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.match_type AS ENUM (
    'preliminary',
    'quarter_final',
    'semi_final',
    'final',
    'eliminator'
);


ALTER TYPE public.match_type OWNER TO postgres;

--
-- Name: playing_format; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.playing_format AS ENUM (
    'test',
    'odi',
    't20i',
    'first_class',
    'list_a',
    't20'
);


ALTER TYPE public.playing_format OWNER TO postgres;

--
-- Name: playing_level; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.playing_level AS ENUM (
    'international',
    'domestic'
);


ALTER TYPE public.playing_level OWNER TO postgres;

--
-- Name: playing_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.playing_status AS ENUM (
    'playing_xi',
    'bench',
    'substitute',
    'withdrawn',
    'impact_player'
);


ALTER TYPE public.playing_status OWNER TO postgres;

--
-- Name: user_role; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.user_role AS ENUM (
    'system_admin',
    'editor_in_chief',
    'editor',
    'sub_editor',
    'scorer'
);


ALTER TYPE public.user_role OWNER TO postgres;

--
-- Name: get_player_profile_by_id(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_player_profile_by_id(player_id integer) RETURNS TABLE(id integer, name text, full_name text, playing_role text, nationality text, is_male boolean, date_of_birth date, image_url text, biography text, is_rhb boolean, bowling_styles public.bowling_style[], primary_bowling_style public.bowling_style, teams_represented jsonb, test_stats jsonb, odi_stats jsonb, t20i_stats jsonb, fc_stats jsonb, lista_stats jsonb, t20_stats jsonb, cricsheet_id text, cricinfo_id text, cricbuzz_id text)
    LANGUAGE plpgsql
    AS $$
 BEGIN RETURN QUERY
SELECT p.id,
    p.name,
	p.full_name,
    p.playing_role,
    p.nationality,
    p.is_male,
    p.date_of_birth,
    p.image_url,
    p.biography,
	p.is_rhb,
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
-- Name: batting_scorecards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.batting_scorecards (
    id integer NOT NULL,
    innings_id integer NOT NULL,
    batter_id integer NOT NULL,
    batting_position integer NOT NULL,
    runs_scored integer DEFAULT 0,
    balls_faced integer DEFAULT 0,
    minutes_batted integer DEFAULT 0,
    fours_scored integer DEFAULT 0,
    sixes_scored integer DEFAULT 0,
    dismissed_by_id integer,
    dismissal_ball_id integer,
    fielder1_id integer,
    fielder2_id integer,
    dismissal_type public.dismissal_type
);


ALTER TABLE public.batting_scorecards OWNER TO postgres;

--
-- Name: batting_scorecards_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.batting_scorecards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.batting_scorecards_id_seq OWNER TO postgres;

--
-- Name: batting_scorecards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.batting_scorecards_id_seq OWNED BY public.batting_scorecards.id;


--
-- Name: blog_articles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.blog_articles (
    id integer NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    author_id integer NOT NULL,
    category public.article_category NOT NULL,
    status public.article_status NOT NULL,
    player_tags integer[],
    team_tags integer[],
    series_tags integer[],
    tournament_tags integer[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.blog_articles OWNER TO postgres;

--
-- Name: blog_articles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.blog_articles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blog_articles_id_seq OWNER TO postgres;

--
-- Name: blog_articles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.blog_articles_id_seq OWNED BY public.blog_articles.id;


--
-- Name: bowling_scorecards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bowling_scorecards (
    id integer NOT NULL,
    innings_id integer NOT NULL,
    bowler_id integer NOT NULL,
    bowling_position integer NOT NULL,
    wickets_taken integer DEFAULT 0,
    runs_conceded integer DEFAULT 0,
    balls_bowled integer DEFAULT 0,
    fours_conceded integer DEFAULT 0,
    sixes_conceded integer DEFAULT 0,
    wides_conceded integer DEFAULT 0,
    noballs_conceded integer DEFAULT 0
);


ALTER TABLE public.bowling_scorecards OWNER TO postgres;

--
-- Name: bowling_scorecards_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bowling_scorecards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bowling_scorecards_id_seq OWNER TO postgres;

--
-- Name: bowling_scorecards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bowling_scorecards_id_seq OWNED BY public.bowling_scorecards.id;


--
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id integer NOT NULL,
    name text NOT NULL,
    host_nation_id integer NOT NULL
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cities_id_seq OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: continents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.continents (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.continents OWNER TO postgres;

--
-- Name: continents_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.continents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.continents_id_seq OWNER TO postgres;

--
-- Name: continents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.continents_id_seq OWNED BY public.continents.id;


--
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
-- Name: grounds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.grounds_id_seq OWNED BY public.grounds.id;


--
-- Name: host_nations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.host_nations (
    id integer NOT NULL,
    name text NOT NULL,
    continent_id integer NOT NULL
);


ALTER TABLE public.host_nations OWNER TO postgres;

--
-- Name: host_nations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.host_nations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.host_nations_id_seq OWNER TO postgres;

--
-- Name: host_nations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.host_nations_id_seq OWNED BY public.host_nations.id;


--
-- Name: innings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.innings (
    id integer NOT NULL,
    match_id integer NOT NULL,
    innings_number integer NOT NULL,
    batting_team_id integer NOT NULL,
    bowling_team_id integer NOT NULL,
    total_runs integer DEFAULT 0,
    total_balls integer DEFAULT 0,
    total_wickets integer DEFAULT 0,
    byes integer DEFAULT 0,
    leg_byes integer DEFAULT 0,
    wides integer DEFAULT 0,
    noballs integer DEFAULT 0,
    penalty integer DEFAULT 0,
    is_super_over boolean DEFAULT false,
    innings_end public.innings_end
);


ALTER TABLE public.innings OWNER TO postgres;

--
-- Name: innings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.innings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.innings_id_seq OWNER TO postgres;

--
-- Name: innings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.innings_id_seq OWNED BY public.innings.id;


--
-- Name: matches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    start_datetime timestamp with time zone,
    team1_id integer,
    team2_id integer,
    is_male boolean NOT NULL,
    tournament_id integer,
    series_id integer,
    host_nation_id integer,
    continent_id integer,
    ground_id integer,
    current_status text,
    home_team_id integer,
    away_team_id integer,
    season text,
    is_day_night boolean,
    is_dls boolean,
    toss_winner_team_id integer,
    toss_loser_team_id integer,
    match_winner_team_id integer,
    match_loser_team_id integer,
    is_won_by_runs boolean,
    win_margin integer,
    balls_remaining_after_win integer,
    potm_id integer,
    scorers_id integer[],
    commentators_id integer[],
    is_toss_decision_bat boolean,
    match_type public.match_type,
    playing_level public.playing_level NOT NULL,
    playing_format public.playing_format NOT NULL,
    tour_id integer,
    final_result public.match_final_result,
    parent_series_id integer,
    balls_per_over integer DEFAULT 6
);


ALTER TABLE public.matches OWNER TO postgres;

--
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
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
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
    full_name text,
    is_rhb boolean,
    bowling_styles public.bowling_style[],
    primary_bowling_style public.bowling_style,
    CONSTRAINT valid_birth_date CHECK ((date_of_birth < CURRENT_DATE)),
    CONSTRAINT valid_image_url CHECK (((image_url IS NULL) OR (image_url ~ '^https?://[^\s/$.?#].[^\s]*$'::text))),
    CONSTRAINT valid_name CHECK ((length(TRIM(BOTH FROM name)) > 0))
);


ALTER TABLE public.players OWNER TO postgres;

--
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
-- Name: players_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.players_id_seq OWNED BY public.players.id;


--
-- Name: seasons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seasons (
    season text NOT NULL
);


ALTER TABLE public.seasons OWNER TO postgres;

--
-- Name: series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series (
    id integer NOT NULL,
    name text NOT NULL,
    is_male boolean NOT NULL,
    playing_level public.playing_level NOT NULL,
    playing_format public.playing_format NOT NULL,
    season text NOT NULL,
    teams_id integer[] NOT NULL,
    host_nations_id integer[] NOT NULL,
    tournament_id integer,
    parent_series_id integer,
    tour_id integer
);


ALTER TABLE public.series OWNER TO postgres;

--
-- Name: series_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.series_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.series_id_seq OWNER TO postgres;

--
-- Name: series_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.series_id_seq OWNED BY public.series.id;


--
-- Name: squads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.squads (
    player_id integer,
    series_id integer,
    match_id integer,
    is_captain boolean DEFAULT false,
    is_wk boolean DEFAULT false,
    is_debut boolean DEFAULT false,
    playing_status public.playing_status
);


ALTER TABLE public.squads OWNER TO postgres;

--
-- Name: teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.teams (
    id integer NOT NULL,
    name text NOT NULL,
    is_male boolean DEFAULT true,
    image_url text,
    short_name text,
    playing_level public.playing_level DEFAULT 'international'::public.playing_level NOT NULL
);


ALTER TABLE public.teams OWNER TO postgres;

--
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
-- Name: teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.teams_id_seq OWNED BY public.teams.id;


--
-- Name: tournaments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tournaments (
    id integer NOT NULL,
    name text NOT NULL,
    is_male boolean NOT NULL,
    playing_level public.playing_level NOT NULL,
    playing_format public.playing_format NOT NULL
);


ALTER TABLE public.tournaments OWNER TO postgres;

--
-- Name: tournaments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tournaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tournaments_id_seq OWNER TO postgres;

--
-- Name: tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;


--
-- Name: tours; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tours (
    id integer NOT NULL,
    "touring_team_Id" integer NOT NULL,
    host_nations_id integer[] NOT NULL,
    season text NOT NULL
);


ALTER TABLE public.tours OWNER TO postgres;

--
-- Name: tours_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tours_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tours_id_seq OWNER TO postgres;

--
-- Name: tours_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tours_id_seq OWNED BY public.tours.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role public.user_role NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: batting_scorecards id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards ALTER COLUMN id SET DEFAULT nextval('public.batting_scorecards_id_seq'::regclass);


--
-- Name: blog_articles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blog_articles ALTER COLUMN id SET DEFAULT nextval('public.blog_articles_id_seq'::regclass);


--
-- Name: bowling_scorecards id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards ALTER COLUMN id SET DEFAULT nextval('public.bowling_scorecards_id_seq'::regclass);


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: continents id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.continents ALTER COLUMN id SET DEFAULT nextval('public.continents_id_seq'::regclass);


--
-- Name: grounds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds ALTER COLUMN id SET DEFAULT nextval('public.grounds_id_seq'::regclass);


--
-- Name: host_nations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.host_nations ALTER COLUMN id SET DEFAULT nextval('public.host_nations_id_seq'::regclass);


--
-- Name: innings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings ALTER COLUMN id SET DEFAULT nextval('public.innings_id_seq'::regclass);


--
-- Name: matches id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches ALTER COLUMN id SET DEFAULT nextval('public.matches_id_seq'::regclass);


--
-- Name: players id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players ALTER COLUMN id SET DEFAULT nextval('public.players_id_seq'::regclass);


--
-- Name: series id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series ALTER COLUMN id SET DEFAULT nextval('public.series_id_seq'::regclass);


--
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- Name: tournaments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);


--
-- Name: tours id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tours ALTER COLUMN id SET DEFAULT nextval('public.tours_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: batting_scorecards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.batting_scorecards (id, innings_id, batter_id, batting_position, runs_scored, balls_faced, minutes_batted, fours_scored, sixes_scored, dismissed_by_id, dismissal_ball_id, fielder1_id, fielder2_id, dismissal_type) FROM stdin;
\.


--
-- Data for Name: blog_articles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.blog_articles (id, title, content, author_id, category, status, player_tags, team_tags, series_tags, tournament_tags, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: bowling_scorecards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bowling_scorecards (id, innings_id, bowler_id, bowling_position, wickets_taken, runs_conceded, balls_bowled, fours_conceded, sixes_conceded, wides_conceded, noballs_conceded) FROM stdin;
\.


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, name, host_nation_id) FROM stdin;
1	Mumbai	1
\.


--
-- Data for Name: continents; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.continents (id, name) FROM stdin;
1	Asia
\.


--
-- Data for Name: grounds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.grounds (id, name, host_nation_id, city_id) FROM stdin;
1	Wankhede Stadium	1	1
\.


--
-- Data for Name: host_nations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.host_nations (id, name, continent_id) FROM stdin;
1	India	1
\.


--
-- Data for Name: innings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.innings (id, match_id, innings_number, batting_team_id, bowling_team_id, total_runs, total_balls, total_wickets, byes, leg_byes, wides, noballs, penalty, is_super_over, innings_end) FROM stdin;
\.


--
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.matches (id, start_datetime, team1_id, team2_id, is_male, tournament_id, series_id, host_nation_id, continent_id, ground_id, current_status, home_team_id, away_team_id, season, is_day_night, is_dls, toss_winner_team_id, toss_loser_team_id, match_winner_team_id, match_loser_team_id, is_won_by_runs, win_margin, balls_remaining_after_win, potm_id, scorers_id, commentators_id, is_toss_decision_bat, match_type, playing_level, playing_format, tour_id, final_result, parent_series_id, balls_per_over) FROM stdin;
\.


--
-- Data for Name: players; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.players (id, name, playing_role, nationality, is_male, date_of_birth, image_url, biography, teams_represented_id, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id, full_name, is_rhb, bowling_styles, primary_bowling_style) FROM stdin;
1	Virat Kohli	Top-order Batter	India	t	1994-11-05	\N	\N	{1,3}	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	\N	\N	\N	Virat Kohli	t	\N	\N
2	Rohit Sharma	Top-order Batter	India	t	1993-04-30	\N	\N	\N	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	\N	\N	\N	Rohit Gurunath Sharma	t	\N	\N
3	Suryakumar Yadav	Batter	India	t	1990-09-14	\N	\N	{1,4}	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	(,,,,,,,,,,,,,,,,,,,,,,,,,)	\N	\N	\N	Suryakumar Ashok Yadav	t	\N	\N
\.


--
-- Data for Name: seasons; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seasons (season) FROM stdin;
2023/24
\.


--
-- Data for Name: series; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.series (id, name, is_male, playing_level, playing_format, season, teams_id, host_nations_id, tournament_id, parent_series_id, tour_id) FROM stdin;
1	ICC Cricket World Cup 2023	t	international	odi	2023/24	{1,2}	{1}	1	\N	\N
\.


--
-- Data for Name: squads; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.squads (player_id, series_id, match_id, is_captain, is_wk, is_debut, playing_status) FROM stdin;
\.


--
-- Data for Name: teams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.teams (id, name, is_male, image_url, short_name, playing_level) FROM stdin;
1	India	t	\N	IND	international
2	Australia	t	\N	AUS	international
3	Royal Challengers Bangalore	t	\N	RCB	domestic
4	Mumbai Indians	t	\N	\N	domestic
\.


--
-- Data for Name: tournaments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tournaments (id, name, is_male, playing_level, playing_format) FROM stdin;
1	ICC Men's World Cup	t	international	odi
\.


--
-- Data for Name: tours; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tours (id, "touring_team_Id", host_nations_id, season) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, role) FROM stdin;
\.


--
-- Name: batting_scorecards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.batting_scorecards_id_seq', 1, false);


--
-- Name: blog_articles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.blog_articles_id_seq', 1, false);


--
-- Name: bowling_scorecards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bowling_scorecards_id_seq', 1, false);


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 1, true);


--
-- Name: continents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.continents_id_seq', 1, true);


--
-- Name: grounds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.grounds_id_seq', 1, true);


--
-- Name: host_nations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.host_nations_id_seq', 1, true);


--
-- Name: innings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.innings_id_seq', 1, false);


--
-- Name: matches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.matches_id_seq', 1, false);


--
-- Name: players_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.players_id_seq', 3, true);


--
-- Name: series_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.series_id_seq', 1, true);


--
-- Name: teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.teams_id_seq', 4, true);


--
-- Name: tournaments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tournaments_id_seq', 1, true);


--
-- Name: tours_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tours_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: batting_scorecards batting_scorecards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_pkey PRIMARY KEY (id);


--
-- Name: blog_articles blog_articles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blog_articles
    ADD CONSTRAINT blog_articles_pkey PRIMARY KEY (id);


--
-- Name: bowling_scorecards bowling_scorecards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_pkey PRIMARY KEY (id);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- Name: continents continents_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.continents
    ADD CONSTRAINT continents_pkey PRIMARY KEY (id);


--
-- Name: grounds grounds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds
    ADD CONSTRAINT grounds_pkey PRIMARY KEY (id);


--
-- Name: host_nations host_nations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.host_nations
    ADD CONSTRAINT host_nations_pkey PRIMARY KEY (id);


--
-- Name: innings innings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_pkey PRIMARY KEY (id);


--
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);


--
-- Name: seasons seasons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seasons
    ADD CONSTRAINT seasons_pkey PRIMARY KEY (season);


--
-- Name: series series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_pkey PRIMARY KEY (id);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (id);


--
-- Name: tournaments tournaments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_pkey PRIMARY KEY (id);


--
-- Name: tours tours_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tours
    ADD CONSTRAINT tours_pkey PRIMARY KEY (id);


--
-- Name: continents unique_continent_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.continents
    ADD CONSTRAINT unique_continent_name UNIQUE (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: batting_scorecards batting_scorecards_batter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_batter_id_fkey FOREIGN KEY (batter_id) REFERENCES public.players(id);


--
-- Name: batting_scorecards batting_scorecards_dismissed_by_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_dismissed_by_id_fkey FOREIGN KEY (dismissed_by_id) REFERENCES public.players(id);


--
-- Name: batting_scorecards batting_scorecards_fielder1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_fielder1_id_fkey FOREIGN KEY (fielder1_id) REFERENCES public.players(id);


--
-- Name: batting_scorecards batting_scorecards_fielder2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_fielder2_id_fkey FOREIGN KEY (fielder2_id) REFERENCES public.players(id);


--
-- Name: batting_scorecards batting_scorecards_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id);


--
-- Name: bowling_scorecards bowling_scorecards_bowler_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_bowler_id_fkey FOREIGN KEY (bowler_id) REFERENCES public.players(id);


--
-- Name: bowling_scorecards bowling_scorecards_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id);


--
-- Name: cities cities_host_nation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_host_nation_id_fkey FOREIGN KEY (host_nation_id) REFERENCES public.host_nations(id);


--
-- Name: grounds grounds_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds
    ADD CONSTRAINT grounds_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) NOT VALID;


--
-- Name: grounds grounds_host_nation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds
    ADD CONSTRAINT grounds_host_nation_id_fkey FOREIGN KEY (host_nation_id) REFERENCES public.host_nations(id) NOT VALID;


--
-- Name: host_nations host_nations_continent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.host_nations
    ADD CONSTRAINT host_nations_continent_id_fkey FOREIGN KEY (continent_id) REFERENCES public.continents(id);


--
-- Name: innings innings_batting_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_batting_team_id_fkey FOREIGN KEY (batting_team_id) REFERENCES public.teams(id);


--
-- Name: innings innings_bowling_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_bowling_team_id_fkey FOREIGN KEY (bowling_team_id) REFERENCES public.teams(id);


--
-- Name: innings innings_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id);


--
-- Name: matches matches_away_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_away_team_id_fkey FOREIGN KEY (away_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_continent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_continent_id_fkey FOREIGN KEY (continent_id) REFERENCES public.continents(id) NOT VALID;


--
-- Name: matches matches_ground_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_ground_id_fkey FOREIGN KEY (ground_id) REFERENCES public.grounds(id) NOT VALID;


--
-- Name: matches matches_home_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_home_team_id_fkey FOREIGN KEY (home_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_host_nation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_host_nation_id_fkey FOREIGN KEY (host_nation_id) REFERENCES public.host_nations(id) NOT VALID;


--
-- Name: matches matches_match_loser_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_match_loser_team_id_fkey FOREIGN KEY (match_loser_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_match_winner_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_match_winner_team_id_fkey FOREIGN KEY (match_winner_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_parent_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_parent_series_id_fkey FOREIGN KEY (parent_series_id) REFERENCES public.series(id) NOT VALID;


--
-- Name: matches matches_potm_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_potm_id_fkey FOREIGN KEY (potm_id) REFERENCES public.players(id) NOT VALID;


--
-- Name: matches matches_season_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_season_fkey FOREIGN KEY (season) REFERENCES public.seasons(season) NOT VALID;


--
-- Name: matches matches_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id) NOT VALID;


--
-- Name: matches matches_team1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team1_id_fkey FOREIGN KEY (team1_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_team2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team2_id_fkey FOREIGN KEY (team2_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_toss_loser_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_toss_loser_team_id_fkey FOREIGN KEY (toss_loser_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_toss_winner_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_toss_winner_team_id_fkey FOREIGN KEY (toss_winner_team_id) REFERENCES public.teams(id) NOT VALID;


--
-- Name: matches matches_tour_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_tour_id_fkey FOREIGN KEY (tour_id) REFERENCES public.tours(id) NOT VALID;


--
-- Name: matches matches_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) NOT VALID;


--
-- Name: series series_parent_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_parent_series_id_fkey FOREIGN KEY (parent_series_id) REFERENCES public.series(id) NOT VALID;


--
-- Name: series series_season_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_season_fkey FOREIGN KEY (season) REFERENCES public.seasons(season);


--
-- Name: series series_tour_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_tour_id_fkey FOREIGN KEY (tour_id) REFERENCES public.tours(id);


--
-- Name: series series_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id);


--
-- Name: squads squads_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.squads
    ADD CONSTRAINT squads_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id);


--
-- Name: squads squads_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.squads
    ADD CONSTRAINT squads_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id);


--
-- Name: squads squads_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.squads
    ADD CONSTRAINT squads_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id);


--
-- Name: tours tours_season_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tours
    ADD CONSTRAINT tours_season_fkey FOREIGN KEY (season) REFERENCES public.seasons(season);


--
-- Name: tours tours_touring_team_Id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tours
    ADD CONSTRAINT "tours_touring_team_Id_fkey" FOREIGN KEY ("touring_team_Id") REFERENCES public.teams(id);


--
-- PostgreSQL database dump complete
--

