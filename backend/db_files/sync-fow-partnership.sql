CREATE OR REPLACE FUNCTION sync_fow_partnership (p_innings_id INT) returns void AS $$
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
$$ language plpgsql;