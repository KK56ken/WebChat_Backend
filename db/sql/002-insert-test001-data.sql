INSERT INTO users (email, name, password, token, avatorurl) VALUES ('test@test.com', 'テスト太郎', 'password123', 'aaa', 'img/user/1');
INSERT INTO users (email, name, password, token, avatorurl) VALUES ('test2@test2.com', 'うらしま', 'password', 'bbb', 'img/user/2');
INSERT INTO users (email, name, password, token, avatorurl) VALUES ('test3@test3.com', 'あらた', 'pass', 'ccc', 'img/user/3');


INSERT INTO messages (sendUserId, receiveUserId, message, time) VALUES (1, 2, 'おはよう',CURRENT_TIMESTAMP);
INSERT INTO messages (sendUserId, receiveUserId, message, time) VALUES (2, 1, 'hello',CURRENT_TIMESTAMP);

INSERT INTO friends (sendFriendId, receiveFriendId) VALUES (1, 2);
INSERT INTO friends (sendFriendId, receiveFriendId) VALUES (1, 3);
INSERT INTO friends (sendFriendId, receiveFriendId) VALUES (2, 1);