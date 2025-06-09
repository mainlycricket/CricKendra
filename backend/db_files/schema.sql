--
-- PostgreSQL database dump
--

-- Dumped from database version 14.18 (Ubuntu 14.18-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.18 (Ubuntu 14.18-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
	not_outs integer,
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
	best_match_fig_wkts integer
);


ALTER TYPE public.career_stats OWNER TO postgres;

--
-- Name: dismissal_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.dismissal_type AS ENUM (
    'caught',
    'bowled',
    'lbw',
    'run out',
    'stumped',
    'hit wicket',
    'handled the ball',
    'obstructing the field',
    'timed out',
    'retired hurt',
    'hit the ball twice',
    'caught and bowled',
    'retired out',
    'retired not out'
);


ALTER TYPE public.dismissal_type OWNER TO postgres;

--
-- Name: innings_end; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.innings_end AS ENUM (
    'all_out',
    'declared',
    'target_reached',
    'forfeited'
);


ALTER TYPE public.innings_end OWNER TO postgres;

--
-- Name: match_final_result; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.match_final_result AS ENUM (
    'winner decided',
    'tie',
    'draw',
    'no result',
    'abandoned'
);


ALTER TYPE public.match_final_result OWNER TO postgres;

--
-- Name: match_state_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.match_state_enum AS ENUM (
    'upcoming',
    'live',
    'break',
    'completed'
);


ALTER TYPE public.match_state_enum OWNER TO postgres;

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
-- Name: outcome_special_method; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.outcome_special_method AS ENUM (
    'D/L',
    'VJD',
    'Awarded',
    '1st innings score',
    'Lost fewer wickets'
);


ALTER TYPE public.outcome_special_method OWNER TO postgres;

--
-- Name: player_award; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.player_award AS ENUM (
    'player_of_the_match',
    'player_of_the_series'
);


ALTER TYPE public.player_award OWNER TO postgres;

--
-- Name: playing_format; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.playing_format AS ENUM (
    'Test',
    'ODI',
    'T20I',
    'first_class',
    'list_a',
    'T20'
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
-- Name: tour_flag; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.tour_flag AS ENUM (
    'tour_series',
    'tour_sub_series'
);


ALTER TYPE public.tour_flag OWNER TO postgres;

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
-- Name: combine_career_stats(public.career_stats, public.career_stats); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.combine_career_stats(stats1 public.career_stats, stats2 public.career_stats) RETURNS public.career_stats
    LANGUAGE plpgsql
    AS $$
DECLARE combined career_stats;
BEGIN -- Initialize the combined stats with regular addition for most fields
combined.matches_played := COALESCE(
    stats1.matches_played + stats2.matches_played,
    stats1.matches_played,
    stats2.matches_played
);
IF combined.matches_played IS NULL 
	THEN RETURN NULL ;
END IF;
combined.innings_batted := COALESCE(
    stats1.innings_batted + stats2.innings_batted,
    stats1.innings_batted,
    stats2.innings_batted
);
combined.runs_scored := COALESCE(
    stats1.runs_scored + stats2.runs_scored,
    stats1.runs_scored,
    stats2.runs_scored
);
combined.not_outs := COALESCE(
    stats1.not_outs + stats2.not_outs,
    stats1.not_outs,
    stats2.not_outs
);
combined.balls_faced := COALESCE(
    stats1.balls_faced + stats2.balls_faced,
    stats1.balls_faced,
    stats2.balls_faced
);
combined.fours_scored := COALESCE(
    stats1.fours_scored + stats2.fours_scored,
    stats1.fours_scored,
    stats2.fours_scored
);
combined.sixes_scored := COALESCE(
    stats1.sixes_scored + stats2.sixes_scored,
    stats1.sixes_scored,
    stats2.sixes_scored
);
combined.centuries_scored := COALESCE(
    stats1.centuries_scored + stats2.centuries_scored,
    stats1.centuries_scored,
    stats2.centuries_scored
);
combined.fifties_scored := COALESCE(
    stats1.fifties_scored + stats2.fifties_scored,
    stats1.fifties_scored,
    stats2.fifties_scored
);
combined.innings_bowled := COALESCE(
    stats1.innings_bowled + stats2.innings_bowled,
    stats1.innings_bowled,
    stats2.innings_bowled
);
combined.runs_conceded := COALESCE(
    stats1.runs_conceded + stats2.runs_conceded,
    stats1.runs_conceded,
    stats2.runs_conceded
);
combined.wickets_taken := COALESCE(
    stats1.wickets_taken + stats2.wickets_taken,
    stats1.wickets_taken,
    stats2.wickets_taken
);
combined.balls_bowled := COALESCE(
    stats1.balls_bowled + stats2.balls_bowled,
    stats1.balls_bowled,
    stats2.balls_bowled
);
combined.fours_conceded := COALESCE(
    stats1.fours_conceded + stats2.fours_conceded,
    stats1.fours_conceded,
    stats2.fours_conceded
);
combined.sixes_conceded := COALESCE(
    stats1.sixes_conceded + stats2.sixes_conceded,
    stats1.sixes_conceded,
    stats2.sixes_conceded
);
combined.four_wkt_hauls := COALESCE(
    stats1.four_wkt_hauls + stats2.four_wkt_hauls,
    stats1.four_wkt_hauls,
    stats2.four_wkt_hauls
);
combined.five_wkt_hauls := COALESCE(
    stats1.five_wkt_hauls + stats2.five_wkt_hauls,
    stats1.five_wkt_hauls,
    stats2.five_wkt_hauls
);
combined.ten_wkt_hauls := COALESCE(
    stats1.ten_wkt_hauls + stats2.ten_wkt_hauls,
    stats1.ten_wkt_hauls,
    stats2.ten_wkt_hauls
);
-- Handle highest score logic
IF stats1.highest_score IS NULL
AND stats2.is_highest_not_out IS NULL THEN combined.highest_score := NULL;
combined.is_highest_not_out := NULL;
ELSIF COALESCE(stats1.highest_score, 0) > COALESCE(stats2.highest_score, 0) THEN combined.highest_score := stats1.highest_score;
combined.is_highest_not_out := stats1.is_highest_not_out;
ELSIF COALESCE(stats1.highest_score, 0) < COALESCE(stats2.highest_score, 0) THEN combined.highest_score := stats2.highest_score;
combined.is_highest_not_out := stats2.is_highest_not_out;
ELSE -- If scores are equal, prefer the not out innings
IF COALESCE(stats1.is_highest_not_out, false)
AND NOT COALESCE(stats2.is_highest_not_out, false) THEN combined.highest_score := stats1.highest_score;
combined.is_highest_not_out := true;
ELSE combined.highest_score := stats2.highest_score;
combined.is_highest_not_out := stats2.is_highest_not_out;
END IF;
END IF;
-- Handle best innings bowling figures
IF stats1.best_inn_fig_wkts is NULL
AND stats2.best_inn_fig_wkts is NULL THEN combined.best_inn_fig_wkts := NULL;
combined.best_inn_fig_runs := NULL;
ELSIF COALESCE(stats1.best_inn_fig_wkts, 0) > COALESCE(stats2.best_inn_fig_wkts, 0) THEN combined.best_inn_fig_wkts := stats1.best_inn_fig_wkts;
combined.best_inn_fig_runs := stats1.best_inn_fig_runs;
ELSIF COALESCE(stats1.best_inn_fig_wkts, 0) < COALESCE(stats2.best_inn_fig_wkts, 0) THEN combined.best_inn_fig_wkts := stats2.best_inn_fig_wkts;
combined.best_inn_fig_runs := stats2.best_inn_fig_runs;
ELSE -- If wickets are equal, prefer the one with fewer runs
IF COALESCE(stats1.best_inn_fig_runs, 999999) < COALESCE(stats2.best_inn_fig_runs, 999999) THEN combined.best_inn_fig_wkts := stats1.best_inn_fig_wkts;
combined.best_inn_fig_runs := stats1.best_inn_fig_runs;
ELSE combined.best_inn_fig_wkts := stats2.best_inn_fig_wkts;
combined.best_inn_fig_runs := stats2.best_inn_fig_runs;
END IF;
END IF;
-- Handle best match bowling figures
IF stats1.best_match_fig_wkts IS NULL
AND stats2.best_match_fig_wkts IS NULL THEN combined.best_match_fig_wkts := NULL;
combined.best_match_fig_runs := NULL;
ELSIF COALESCE(stats1.best_match_fig_wkts, 0) > COALESCE(stats2.best_match_fig_wkts, 0) THEN combined.best_match_fig_wkts := stats1.best_match_fig_wkts;
combined.best_match_fig_runs := stats1.best_match_fig_runs;
ELSIF COALESCE(stats1.best_match_fig_wkts, 0) < COALESCE(stats2.best_match_fig_wkts, 0) THEN combined.best_match_fig_wkts := stats2.best_match_fig_wkts;
combined.best_match_fig_runs := stats2.best_match_fig_runs;
ELSE -- If wickets are equal, prefer the one with fewer runs
IF COALESCE(stats1.best_match_fig_runs, 999999) < COALESCE(stats2.best_match_fig_runs, 999999) THEN combined.best_match_fig_wkts := stats1.best_match_fig_wkts;
combined.best_match_fig_runs := stats1.best_match_fig_runs;
ELSE combined.best_match_fig_wkts := stats2.best_match_fig_wkts;
combined.best_match_fig_runs := stats2.best_match_fig_runs;
END IF;
END IF;
RETURN combined;
END;
$$;


ALTER FUNCTION public.combine_career_stats(stats1 public.career_stats, stats2 public.career_stats) OWNER TO postgres;

--
-- Name: sync_fow_partnership(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sync_fow_partnership(p_innings_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
	delivery_record RECORD;
	TOTAL_RUN INTEGER := 0;
	TOTAL_WKTS INTEGER := 0;
	BATTER1_ID INTEGER := 0;
	BATTER2_ID INTEGER := 0;
	BATTER1_RUN INTEGER := 0;
	BATTER2_RUN INTEGER := 0;
	BATTER1_BALL INTEGER := 0;
	BATTER2_BALL INTEGER := 0;
	START_PARTNERSHIP BOOLEAN := TRUE;
BEGIN
	DELETE FROM fall_of_wickets WHERE innings_id = p_innings_id;
	DELETE FROM batting_partnerships WHERE innings_id = p_innings_id;

	FOR delivery_record IN
		SELECT batter_id, non_striker_id, batter_runs, wides, total_runs, ball_number, innings_delivery_number, 
			player1_dismissed_id, player1_dismissal_type, player2_dismissed_id, player2_dismissal_type
		FROM deliveries
		WHERE innings_id = p_innings_id ORDER BY innings_delivery_number
		LOOP
			IF START_PARTNERSHIP THEN
				BATTER1_ID := delivery_record.batter_id;
				BATTER2_ID := delivery_record.non_striker_id;
				BATTER1_RUN := 0;
				BATTER2_RUN := 0;
				BATTER1_BALL := 0;
				BATTER2_BALL := 0;
				INSERT INTO batting_partnerships(innings_id, wicket_number, start_innings_delivery_number, end_innings_delivery_number, 
					batter1_id, batter1_runs, batter1_balls, batter2_id, batter2_runs, batter2_balls, 
					start_team_runs, end_team_runs, start_ball_number, end_ball_number, is_unbeaten
				) VALUES (p_innings_id, TOTAL_WKTS + 1, delivery_record.innings_delivery_number, 
				delivery_record.innings_delivery_number, delivery_record.batter_id, 0, 0, 
				delivery_record.non_striker_id, 0, 0,
					TOTAL_RUN, TOTAL_RUN, delivery_record.ball_number, delivery_record.ball_number, TRUE 
				);
				START_PARTNERSHIP := FALSE;
			END IF;
			
			IF delivery_record.batter_id = BATTER1_ID THEN
				BATTER1_RUN := BATTER1_RUN + delivery_record.batter_runs;
				IF delivery_record.wides = 0 THEN
					BATTER1_BALL := BATTER1_BALL + 1;
				END IF;
			ELSE
				BATTER2_RUN := BATTER2_RUN + delivery_record.batter_runs;
				IF delivery_record.wides = 0 THEN
					BATTER2_BALL := BATTER2_BALL + 1;
				END IF;
			END IF;
		
			TOTAL_RUN := TOTAL_RUN + delivery_record.total_runs;
			IF delivery_record.player1_dismissed_id IS NOT NULL THEN
				IF delivery_record.player1_dismissal_type NOT IN ('retired hurt', 'retired not out') THEN
					TOTAL_WKTS := TOTAL_WKTS + 1;
				END IF;
				INSERT INTO fall_of_wickets (innings_id, batter_id, team_runs, wicket_number, 
					dismissal_type, innings_delivery_number, ball_number
					) VALUES (p_innings_id, delivery_record.player1_dismissed_id, TOTAL_RUN, TOTAL_WKTS,
					delivery_record.player1_dismissal_type, delivery_record.innings_delivery_number, delivery_record.ball_number);
				UPDATE batting_partnerships SET
					end_innings_delivery_number = delivery_record.innings_delivery_number, is_unbeaten = FALSE,
					batter1_runs = BATTER1_RUN, batter1_balls = BATTER1_BALL,
					batter2_runs = BATTER2_RUN, batter2_balls = BATTER2_BALL,
					end_team_runs = TOTAL_RUN, end_ball_number = delivery_record.ball_number
				WHERE innings_id = p_innings_id AND is_unbeaten = TRUE;

				START_PARTNERSHIP := TRUE;
			END IF;
			IF delivery_record.player2_dismissed_id IS NOT NULL THEN
				IF delivery_record.player2_dismissal_type NOT IN ('retired hurt', 'retired not out') THEN
 					TOTAL_WKTS := TOTAL_WKTS + 1;
				END IF;

				IF delivery_record.player2_dismissed_id != delivery_record.batter_id AND delivery_record.player2_dismissed_id != delivery_record.non_striker_id THEN
					BATTER1_ID := delivery_record.player2_dismissed_id;
					IF delivery_record.player1_dismissed_id = delivery_record.batter_id THEN
						BATTER2_ID := delivery_record.non_striker_id;
					ELSE
						BATTER2_ID := delivery_record.batter_id;
					END IF;

					INSERT INTO batting_partnerships(innings_id, wicket_number, start_innings_delivery_number, end_innings_delivery_number, 
					batter1_id, batter1_runs, batter1_balls, batter2_id, batter2_runs, batter2_balls, 
					start_team_runs, end_team_runs, start_ball_number, end_ball_number, is_unbeaten
					) VALUES (p_innings_id, TOTAL_WKTS - 1, delivery_record.innings_delivery_number, 
					delivery_record.innings_delivery_number, BATTER1_ID, 0, 0, 
					BATTER2_ID, 0, 0,
						TOTAL_RUN, TOTAL_RUN, delivery_record.ball_number, delivery_record.ball_number, FALSE 
					);
					
					START_PARTNERSHIP := TRUE;
				END IF;

				INSERT INTO fall_of_wickets (innings_id, batter_id, team_runs, wicket_number, 
					dismissal_type, innings_delivery_number, ball_number
					) VALUES (p_innings_id, delivery_record.player2_dismissed_id, TOTAL_RUN, TOTAL_WKTS,
					delivery_record.player2_dismissal_type, delivery_record.innings_delivery_number, delivery_record.ball_number);
			END IF;

		END LOOP;

		IF START_PARTNERSHIP = FALSE THEN
			WITH deliveries_data AS (
				SELECT MAX(innings_delivery_number) AS max_delivery_number,
					MAX(ball_number) AS max_ball_number
				FROM deliveries
				WHERE innings_id = p_innings_id
			)
			UPDATE batting_partnerships SET
				end_innings_delivery_number = deliveries_data.max_delivery_number,
				batter1_runs = BATTER1_RUN, batter1_balls = BATTER1_BALL,
				batter2_runs = BATTER2_RUN, batter2_balls = BATTER2_BALL,
				end_team_runs = TOTAL_RUN, end_ball_number = deliveries_data.max_ball_number
			FROM deliveries_data
			WHERE innings_id = p_innings_id AND is_unbeaten = TRUE;
		END IF;
END;
$$;


ALTER FUNCTION public.sync_fow_partnership(p_innings_id integer) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: batting_partnerships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.batting_partnerships (
    innings_id integer NOT NULL,
    wicket_number integer NOT NULL,
    start_innings_delivery_number integer NOT NULL,
    end_innings_delivery_number integer NOT NULL,
    batter1_id integer NOT NULL,
    batter1_runs integer DEFAULT 0 NOT NULL,
    batter1_balls integer DEFAULT 0 NOT NULL,
    batter2_id integer NOT NULL,
    batter2_runs integer DEFAULT 0 NOT NULL,
    batter2_balls integer DEFAULT 0 NOT NULL,
    start_team_runs integer NOT NULL,
    end_team_runs integer NOT NULL,
    start_ball_number double precision NOT NULL,
    end_ball_number double precision NOT NULL,
    is_unbeaten boolean DEFAULT true NOT NULL
);


ALTER TABLE public.batting_partnerships OWNER TO postgres;

--
-- Name: batting_scorecards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.batting_scorecards (
    innings_id integer NOT NULL,
    batter_id integer NOT NULL,
    batting_position integer,
    runs_scored integer DEFAULT 0,
    balls_faced integer DEFAULT 0,
    minutes_batted integer DEFAULT 0,
    fours_scored integer DEFAULT 0,
    sixes_scored integer DEFAULT 0,
    dismissed_by_id integer,
    fielder1_id integer,
    fielder2_id integer,
    dismissal_type public.dismissal_type,
    has_batted boolean DEFAULT false,
    dismissal_ball_number integer
);


ALTER TABLE public.batting_scorecards OWNER TO postgres;

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


ALTER TABLE public.blog_articles_id_seq OWNER TO postgres;

--
-- Name: blog_articles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.blog_articles_id_seq OWNED BY public.blog_articles.id;


--
-- Name: bowling_scorecards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bowling_scorecards (
    innings_id integer NOT NULL,
    bowler_id integer NOT NULL,
    bowling_position integer,
    wickets_taken integer DEFAULT 0,
    runs_conceded integer DEFAULT 0,
    balls_bowled integer DEFAULT 0,
    fours_conceded integer DEFAULT 0,
    sixes_conceded integer DEFAULT 0,
    wides_conceded integer DEFAULT 0,
    noballs_conceded integer DEFAULT 0,
    maiden_overs integer DEFAULT 0
);


ALTER TABLE public.bowling_scorecards OWNER TO postgres;

--
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id integer NOT NULL,
    name text NOT NULL,
    host_nation_id integer
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


ALTER TABLE public.cities_id_seq OWNER TO postgres;

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


ALTER TABLE public.continents_id_seq OWNER TO postgres;

--
-- Name: continents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.continents_id_seq OWNED BY public.continents.id;


--
-- Name: cricsheet_people; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cricsheet_people (
    identifier text NOT NULL,
    name text NOT NULL,
    unique_name text,
    key_bcci text,
    key_bcci_2 text,
    key_bigbash text,
    key_cricbuzz text,
    key_cricheroes text,
    key_crichq text,
    key_cricinfo text,
    key_cricinfo_2 text,
    key_cricingif text,
    key_cricketarchive text,
    key_cricketarchive_2 text,
    key_cricketworld text,
    key_nvplay text,
    key_nvplay_2 text,
    key_opta text,
    key_opta_2 text,
    key_pulse text,
    key_pulse_2 text
);


ALTER TABLE public.cricsheet_people OWNER TO postgres;

--
-- Name: deliveries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.deliveries (
    innings_id integer NOT NULL,
    ball_number double precision,
    over_number integer,
    batter_id integer,
    bowler_id integer,
    non_striker_id integer,
    batter_runs integer,
    wides integer,
    noballs integer,
    legbyes integer,
    byes integer,
    penalty integer,
    total_runs integer,
    bowler_runs integer,
    is_four boolean,
    is_six boolean,
    player1_dismissed_id integer,
    player1_dismissal_type public.dismissal_type,
    player2_dismissed_id integer,
    player2_dismissal_type public.dismissal_type,
    is_pace boolean,
    bowling_style public.bowling_style,
    is_batter_rhb boolean,
    is_non_striker_rhb boolean,
    line text,
    length text,
    ball_type text,
    misc text,
    ww_region text,
    foot_type text,
    shot_type text,
    fielder1_id integer,
    fielder2_id integer,
    commentary text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    total_extras integer,
    innings_delivery_number integer NOT NULL,
    ball_speed double precision
);


ALTER TABLE public.deliveries OWNER TO postgres;

--
-- Name: fall_of_wickets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fall_of_wickets (
    innings_id integer NOT NULL,
    batter_id integer NOT NULL,
    team_runs integer NOT NULL,
    wicket_number integer NOT NULL,
    dismissal_type public.dismissal_type NOT NULL,
    innings_delivery_number integer NOT NULL,
    ball_number double precision NOT NULL
);


ALTER TABLE public.fall_of_wickets OWNER TO postgres;

--
-- Name: grounds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.grounds (
    id integer NOT NULL,
    name text NOT NULL,
    city_id integer
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


ALTER TABLE public.grounds_id_seq OWNER TO postgres;

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
    continent_id integer
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


ALTER TABLE public.host_nations_id_seq OWNER TO postgres;

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
    innings_number integer,
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
    innings_end public.innings_end,
    target_runs integer,
    max_overs double precision,
    striker_id integer,
    non_striker_id integer,
    bowler1_id integer,
    bowler2_id integer
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


ALTER TABLE public.innings_id_seq OWNER TO postgres;

--
-- Name: innings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.innings_id_seq OWNED BY public.innings.id;


--
-- Name: match_series_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.match_series_entries (
    match_id integer NOT NULL,
    series_id integer NOT NULL
);


ALTER TABLE public.match_series_entries OWNER TO postgres;

--
-- Name: match_squad_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.match_squad_entries (
    player_id integer NOT NULL,
    match_id integer NOT NULL,
    is_captain boolean,
    is_wk boolean,
    is_debut boolean,
    playing_status public.playing_status,
    team_id integer,
    is_vice_captain boolean
);


ALTER TABLE public.match_squad_entries OWNER TO postgres;

--
-- Name: matches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    team1_id integer,
    team2_id integer,
    is_male boolean NOT NULL,
    main_series_id integer,
    ground_id integer,
    current_status text,
    home_team_id integer,
    away_team_id integer,
    season text,
    is_day_night boolean,
    toss_winner_team_id integer,
    toss_loser_team_id integer,
    match_winner_team_id integer,
    match_loser_team_id integer,
    is_won_by_runs boolean,
    win_margin integer,
    balls_remaining_after_win integer,
    is_toss_decision_bat boolean,
    match_type public.match_type,
    playing_level public.playing_level NOT NULL,
    playing_format public.playing_format NOT NULL,
    final_result public.match_final_result,
    balls_per_over integer DEFAULT 6,
    event_match_number integer,
    start_date date,
    bowl_out_winner_id integer,
    super_over_winner_id integer,
    is_won_by_innings boolean,
    outcome_special_method public.outcome_special_method,
    cricsheet_id text,
    is_neutral_venue boolean,
    is_bbb_done boolean DEFAULT false,
    end_date date,
    match_state public.match_state_enum DEFAULT 'upcoming'::public.match_state_enum,
    match_state_description text,
    start_datetime_utc timestamp with time zone
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


ALTER TABLE public.matches_id_seq OWNER TO postgres;

--
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
-- Name: player_awards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player_awards (
    player_id integer NOT NULL,
    match_id integer NOT NULL,
    series_id integer,
    award_type public.player_award NOT NULL
);


ALTER TABLE public.player_awards OWNER TO postgres;

--
-- Name: player_team_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player_team_entries (
    player_id integer NOT NULL,
    team_id integer NOT NULL
);


ALTER TABLE public.player_team_entries OWNER TO postgres;

--
-- Name: players; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.players (
    id integer NOT NULL,
    name text NOT NULL,
    playing_role text,
    nationality text,
    is_male boolean NOT NULL,
    date_of_birth date,
    image_url text,
    biography text,
    db_test_stats public.career_stats,
    db_odi_stats public.career_stats,
    db_t20i_stats public.career_stats,
    db_fc_stats public.career_stats,
    db_lista_stats public.career_stats,
    db_t20_stats public.career_stats,
    cricsheet_id text,
    cricinfo_id text,
    cricbuzz_id text,
    full_name text,
    is_rhb boolean,
    bowling_styles public.bowling_style[],
    primary_bowling_style public.bowling_style,
    unavailable_test_stats public.career_stats,
    unavailable_odi_stats public.career_stats,
    unavailable_t20i_stats public.career_stats,
    unavailable_fc_stats public.career_stats,
    unavailable_lista_stats public.career_stats,
    unavailable_t20_stats public.career_stats,
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


ALTER TABLE public.players_id_seq OWNER TO postgres;

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
    season text,
    tournament_id integer,
    start_date date,
    end_date date,
    winner_team_id integer,
    final_status text,
    tour_flag public.tour_flag
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


ALTER TABLE public.series_id_seq OWNER TO postgres;

--
-- Name: series_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.series_id_seq OWNED BY public.series.id;


--
-- Name: series_squad_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series_squad_entries (
    squad_id integer NOT NULL,
    player_id integer NOT NULL,
    is_captain boolean,
    is_vice_captain boolean,
    is_wk boolean,
    playing_status public.playing_status
);


ALTER TABLE public.series_squad_entries OWNER TO postgres;

--
-- Name: series_squads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series_squads (
    id integer NOT NULL,
    series_id integer NOT NULL,
    team_id integer NOT NULL,
    squad_label text NOT NULL
);


ALTER TABLE public.series_squads OWNER TO postgres;

--
-- Name: series_squads_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.series_squads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.series_squads_id_seq OWNER TO postgres;

--
-- Name: series_squads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.series_squads_id_seq OWNED BY public.series_squads.id;


--
-- Name: series_team_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series_team_entries (
    series_id integer NOT NULL,
    team_id integer NOT NULL
);


ALTER TABLE public.series_team_entries OWNER TO postgres;

--
-- Name: teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.teams (
    id integer NOT NULL,
    name text NOT NULL,
    is_male boolean DEFAULT true NOT NULL,
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


ALTER TABLE public.teams_id_seq OWNER TO postgres;

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


ALTER TABLE public.tournaments_id_seq OWNER TO postgres;

--
-- Name: tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;


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


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: blog_articles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blog_articles ALTER COLUMN id SET DEFAULT nextval('public.blog_articles_id_seq'::regclass);


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
-- Name: series_squads id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squads ALTER COLUMN id SET DEFAULT nextval('public.series_squads_id_seq'::regclass);


--
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- Name: tournaments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: batting_scorecards batting_scorecards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_pkey PRIMARY KEY (innings_id, batter_id);


--
-- Name: blog_articles blog_articles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blog_articles
    ADD CONSTRAINT blog_articles_pkey PRIMARY KEY (id);


--
-- Name: bowling_scorecards bowling_scorecards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_pkey PRIMARY KEY (innings_id, bowler_id);


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
-- Name: cricsheet_people cricsheet_people_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cricsheet_people
    ADD CONSTRAINT cricsheet_people_pkey PRIMARY KEY (identifier);


--
-- Name: deliveries deliveries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_pkey PRIMARY KEY (innings_id, innings_delivery_number);


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
-- Name: innings innings_match_id_innings_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_match_id_innings_number_key UNIQUE (match_id, innings_number);


--
-- Name: innings innings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_pkey PRIMARY KEY (id);


--
-- Name: match_series_entries match_series_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_series_entries
    ADD CONSTRAINT match_series_entries_pkey PRIMARY KEY (match_id, series_id);


--
-- Name: match_squad_entries match_squad_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_squad_entries
    ADD CONSTRAINT match_squad_entries_pkey PRIMARY KEY (player_id, match_id);


--
-- Name: match_squad_entries match_squads_player_id_match_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_squad_entries
    ADD CONSTRAINT match_squads_player_id_match_id_key UNIQUE (player_id, match_id);


--
-- Name: matches matches_cricsheet_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_cricsheet_id_key UNIQUE (cricsheet_id);


--
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- Name: player_awards player_awards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_awards
    ADD CONSTRAINT player_awards_pkey PRIMARY KEY (player_id, match_id, award_type);


--
-- Name: player_team_entries player_team_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_team_entries
    ADD CONSTRAINT player_team_entries_pkey PRIMARY KEY (player_id, team_id);


--
-- Name: player_team_entries player_teams_player_id_team_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_team_entries
    ADD CONSTRAINT player_teams_player_id_team_id_key UNIQUE (player_id, team_id);


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
-- Name: series_squad_entries series_squad_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squad_entries
    ADD CONSTRAINT series_squad_entries_pkey PRIMARY KEY (squad_id, player_id);


--
-- Name: series_squad_entries series_squad_entries_player_id_squad_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squad_entries
    ADD CONSTRAINT series_squad_entries_player_id_squad_id_key UNIQUE (player_id, squad_id);


--
-- Name: series_squads series_squads_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squads
    ADD CONSTRAINT series_squads_pkey PRIMARY KEY (id);


--
-- Name: series_team_entries series_team_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_team_entries
    ADD CONSTRAINT series_team_entries_pkey PRIMARY KEY (series_id, team_id);


--
-- Name: series_team_entries series_team_entries_series_id_team_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_team_entries
    ADD CONSTRAINT series_team_entries_series_id_team_id_key UNIQUE (series_id, team_id);


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
-- Name: continents unique_continent_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.continents
    ADD CONSTRAINT unique_continent_name UNIQUE (name);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: deliveries_innings_id_innings_delivery_number_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX deliveries_innings_id_innings_delivery_number_key ON public.deliveries USING btree (innings_id, innings_delivery_number);


--
-- Name: idx_batting_scorecards_innings_id_batter_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_batting_scorecards_innings_id_batter_id ON public.batting_scorecards USING btree (innings_id, batter_id);


--
-- Name: idx_bowling_scorecards_innings_id_bowler_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_bowling_scorecards_innings_id_bowler_id ON public.bowling_scorecards USING btree (innings_id, bowler_id);


--
-- Name: unique_innings_batter_non_retired; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX unique_innings_batter_non_retired ON public.fall_of_wickets USING btree (innings_id, batter_id) WHERE (dismissal_type <> ALL (ARRAY['retired hurt'::public.dismissal_type, 'retired not out'::public.dismissal_type]));


--
-- Name: unique_innings_wicket_number_non_retired; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX unique_innings_wicket_number_non_retired ON public.fall_of_wickets USING btree (innings_id, wicket_number) WHERE (dismissal_type <> ALL (ARRAY['retired hurt'::public.dismissal_type, 'retired not out'::public.dismissal_type]));


--
-- Name: batting_partnerships batting_partnerships_batter1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_batter1_id_fkey FOREIGN KEY (batter1_id) REFERENCES public.players(id);


--
-- Name: batting_partnerships batting_partnerships_batter2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_batter2_id_fkey FOREIGN KEY (batter2_id) REFERENCES public.players(id);


--
-- Name: batting_partnerships batting_partnerships_innings_id_batter1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_innings_id_batter1_id_fkey FOREIGN KEY (innings_id, batter1_id) REFERENCES public.batting_scorecards(innings_id, batter_id);


--
-- Name: batting_partnerships batting_partnerships_innings_id_batter2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_innings_id_batter2_id_fkey FOREIGN KEY (innings_id, batter2_id) REFERENCES public.batting_scorecards(innings_id, batter_id);


--
-- Name: batting_partnerships batting_partnerships_innings_id_end_innings_delivery_numbe_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_innings_id_end_innings_delivery_numbe_fkey FOREIGN KEY (innings_id, end_innings_delivery_number) REFERENCES public.deliveries(innings_id, innings_delivery_number);


--
-- Name: batting_partnerships batting_partnerships_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id);


--
-- Name: batting_partnerships batting_partnerships_innings_id_start_innings_delivery_num_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_partnerships
    ADD CONSTRAINT batting_partnerships_innings_id_start_innings_delivery_num_fkey FOREIGN KEY (innings_id, start_innings_delivery_number) REFERENCES public.deliveries(innings_id, innings_delivery_number);


--
-- Name: batting_scorecards batting_scorecards_batter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_batter_id_fkey FOREIGN KEY (batter_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: batting_scorecards batting_scorecards_dismissal_ball_number_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_dismissal_ball_number_innings_id_fkey FOREIGN KEY (dismissal_ball_number, innings_id) REFERENCES public.deliveries(innings_delivery_number, innings_id) NOT VALID;


--
-- Name: batting_scorecards batting_scorecards_dismissed_by_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_dismissed_by_id_fkey FOREIGN KEY (dismissed_by_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: batting_scorecards batting_scorecards_fielder1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_fielder1_id_fkey FOREIGN KEY (fielder1_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: batting_scorecards batting_scorecards_fielder2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_fielder2_id_fkey FOREIGN KEY (fielder2_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: batting_scorecards batting_scorecards_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.batting_scorecards
    ADD CONSTRAINT batting_scorecards_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: bowling_scorecards bowling_scorecards_bowler_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_bowler_id_fkey FOREIGN KEY (bowler_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: bowling_scorecards bowling_scorecards_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bowling_scorecards
    ADD CONSTRAINT bowling_scorecards_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: cities cities_host_nation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_host_nation_id_fkey FOREIGN KEY (host_nation_id) REFERENCES public.host_nations(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_batter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_batter_id_fkey FOREIGN KEY (batter_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_bowler_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_bowler_id_fkey FOREIGN KEY (bowler_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_fielder1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_fielder1_id_fkey FOREIGN KEY (fielder1_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_fielder2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_fielder2_id_fkey FOREIGN KEY (fielder2_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_non_striker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_non_striker_id_fkey FOREIGN KEY (non_striker_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_player1_dismissed_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_player1_dismissed_id_fkey FOREIGN KEY (player1_dismissed_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: deliveries deliveries_player2_dismissed_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.deliveries
    ADD CONSTRAINT deliveries_player2_dismissed_id_fkey FOREIGN KEY (player2_dismissed_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: fall_of_wickets fall_of_wickets_batter_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fall_of_wickets
    ADD CONSTRAINT fall_of_wickets_batter_id_fkey FOREIGN KEY (batter_id) REFERENCES public.players(id);


--
-- Name: fall_of_wickets fall_of_wickets_innings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fall_of_wickets
    ADD CONSTRAINT fall_of_wickets_innings_id_fkey FOREIGN KEY (innings_id) REFERENCES public.innings(id);


--
-- Name: grounds grounds_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.grounds
    ADD CONSTRAINT grounds_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: host_nations host_nations_continent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.host_nations
    ADD CONSTRAINT host_nations_continent_id_fkey FOREIGN KEY (continent_id) REFERENCES public.continents(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: innings innings_batting_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_batting_team_id_fkey FOREIGN KEY (batting_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: innings innings_bowler1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_bowler1_id_fkey FOREIGN KEY (bowler1_id) REFERENCES public.players(id);


--
-- Name: innings innings_bowler2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_bowler2_id_fkey FOREIGN KEY (bowler2_id) REFERENCES public.players(id);


--
-- Name: innings innings_bowling_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_bowling_team_id_fkey FOREIGN KEY (bowling_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: fall_of_wickets innings_id_innings_delivery_number_key; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fall_of_wickets
    ADD CONSTRAINT innings_id_innings_delivery_number_key FOREIGN KEY (innings_id, innings_delivery_number) REFERENCES public.deliveries(innings_id, innings_delivery_number);


--
-- Name: innings innings_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: innings innings_non_striker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_non_striker_id_fkey FOREIGN KEY (non_striker_id) REFERENCES public.players(id);


--
-- Name: innings innings_striker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.innings
    ADD CONSTRAINT innings_striker_id_fkey FOREIGN KEY (striker_id) REFERENCES public.players(id);


--
-- Name: match_series_entries match_series_entries_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_series_entries
    ADD CONSTRAINT match_series_entries_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id);


--
-- Name: match_series_entries match_series_entries_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_series_entries
    ADD CONSTRAINT match_series_entries_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id);


--
-- Name: match_squad_entries match_squads_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_squad_entries
    ADD CONSTRAINT match_squads_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_away_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_away_team_id_fkey FOREIGN KEY (away_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_bowl_out_winner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_bowl_out_winner_id_fkey FOREIGN KEY (bowl_out_winner_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_ground_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_ground_id_fkey FOREIGN KEY (ground_id) REFERENCES public.grounds(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_home_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_home_team_id_fkey FOREIGN KEY (home_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_match_loser_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_match_loser_team_id_fkey FOREIGN KEY (match_loser_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_match_winner_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_match_winner_team_id_fkey FOREIGN KEY (match_winner_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_season_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_season_fkey FOREIGN KEY (season) REFERENCES public.seasons(season) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_series_id_fkey FOREIGN KEY (main_series_id) REFERENCES public.series(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_super_over_winner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_super_over_winner_id_fkey FOREIGN KEY (super_over_winner_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_team1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team1_id_fkey FOREIGN KEY (team1_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_team2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team2_id_fkey FOREIGN KEY (team2_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_toss_loser_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_toss_loser_team_id_fkey FOREIGN KEY (toss_loser_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: matches matches_toss_winner_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_toss_winner_team_id_fkey FOREIGN KEY (toss_winner_team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: player_awards player_awards_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_awards
    ADD CONSTRAINT player_awards_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_awards player_awards_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_awards
    ADD CONSTRAINT player_awards_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_awards player_awards_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_awards
    ADD CONSTRAINT player_awards_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_team_entries player_teams_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_team_entries
    ADD CONSTRAINT player_teams_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: player_team_entries player_teams_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_team_entries
    ADD CONSTRAINT player_teams_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: players players_cricsheet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_cricsheet_id_fkey FOREIGN KEY (cricsheet_id) REFERENCES public.cricsheet_people(identifier) NOT VALID;


--
-- Name: series series_season_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_season_fkey FOREIGN KEY (season) REFERENCES public.seasons(season) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: series_squad_entries series_squad_entries_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squad_entries
    ADD CONSTRAINT series_squad_entries_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: series_squad_entries series_squad_entries_squad_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squad_entries
    ADD CONSTRAINT series_squad_entries_squad_id_fkey FOREIGN KEY (squad_id) REFERENCES public.series_squads(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: series_squads series_squads_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squads
    ADD CONSTRAINT series_squads_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: series_squads series_squads_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_squads
    ADD CONSTRAINT series_squads_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: series_team_entries series_team_entries_series_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_team_entries
    ADD CONSTRAINT series_team_entries_series_id_fkey FOREIGN KEY (series_id) REFERENCES public.series(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: series_team_entries series_team_entries_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_team_entries
    ADD CONSTRAINT series_team_entries_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: series series_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: series series_winner_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_winner_team_id_fkey FOREIGN KEY (winner_team_id) REFERENCES public.teams(id);


--
-- Name: match_squad_entries squads_match_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_squad_entries
    ADD CONSTRAINT squads_match_id_fkey FOREIGN KEY (match_id) REFERENCES public.matches(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: match_squad_entries squads_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.match_squad_entries
    ADD CONSTRAINT squads_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- PostgreSQL database dump complete
--

