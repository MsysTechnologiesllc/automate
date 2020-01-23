BEGIN;

INSERT INTO policies
    VALUES ('11fe8e43-c3a6-48dd-9129-4f694e85f370',
            '{"action": "*", "effect": "allow", "resource": "infraProxy:*", "subjects": ["token:*", "user:*"]}',
            CURRENT_TIMESTAMP,
            1,
            TRUE)
    ON CONFLICT (id) DO UPDATE
    SET policy_data='{"action": "*", "effect": "allow", "resource": "infraProxy:*", "subjects": ["token:*", "user:*"]}',
        deletable=TRUE;
END;
