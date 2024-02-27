import React, { useState } from "react";
import ErrorToast from "../../helpers/toast/errorToast";
import searchUsers from "../../helpers/api/webApi/searchUsers/searchForUser";

const SearchUser = ({ setUser }) => {
	const [query, setQuery] = useState("");
	const [users, setUsers] = useState([]);

	const handleSearch = () => {
		searchUsers(query).then((data) => {
			setUsers(data.users);
		}).catch((error) => {
			ErrorToast(error);
		});
	};
	return (
		<div className="w-full min-w-96">
			<div className="">
				<input
					type="text"
					className="w-full p-2 border-2 outline-none rounded-3xl mb-2"
					placeholder="Search user by email"
					value={query}
					onChange={(e) => {
						setQuery(e.target.value);
						handleSearch();
					}}
				/>
				<div className="flex flex-col w-full">
					{users.map((user, i) => {
						return (
							<button
								type="button"
								aria-label={`add ${user.first_name} ${user.last_name} to collaborators`}
								className={`w-full p-3 text-left hover:underline hover:bg-gray-200
											transition-all duration-300
											${i % 2 === 0 ? "bg-gray-100" : ""}`}
								key={user.id}
								onClick={() => {
									setUser(user);
								}}
							>
								{user.first_name} {user.last_name} ({user.email})
							</button>
						);
					})}
				</div>
			</div>
		</div>
	);
};

export default SearchUser;
