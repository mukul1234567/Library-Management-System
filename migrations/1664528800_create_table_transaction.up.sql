
CREATE TABLE `Transactions`(
		`id` VARCHAR(40) NOT NULL,
		`issuedate` VARCHAR(30) NOT NULL,
		`returndate` VARCHAR(30) NOT NULL,
		`duedate` VARCHAR(30) NOT NULL,
		`book_id` VARCHAR(40) NOT NULL,
		`user_id` VARCHAR(40) NOT NULL,
		PRIMARY KEY(id),
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(book_id) REFERENCES books(id) ON DELETE CASCADE
		);
