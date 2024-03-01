import React, { useState } from "react";
import ErrorToast from "../../helpers/toast/errorToast";
import searchGroups from "../../helpers/api/webApi/search/searchForGroup";

const SearchGroup = ({ setGroup }) => {
	const [query, setQuery] = useState("");
	const [groups, setGroups] = useState([]);

	const handleSearch = () => {
		searchGroups(query).then((data) => {
			setGroups(data.groups);
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
					placeholder="Search group"
					value={query}
					onChange={(e) => {
						setQuery(e.target.value);
						handleSearch();
					}}
				/>
				<div className="flex flex-col w-full">
					{groups.map((group, i) => {
						return (
							<button
								type="button"
								aria-label={`add ${group.group_name} to collaborators`}
								className={`w-full p-3 text-left hover:underline hover:bg-gray-200
											transition-all duration-300
											${i % 2 === 0 ? "bg-gray-100" : ""}`}
								key={group.id}
								onClick={() => {
									setQuery(`${group.group_name}`);
									setGroups([]);
									setGroup(group);
								}}
							>
								{group.group_name}
							</button>
						);
					})}
				</div>
			</div>
		</div>
	);
};

export default SearchGroup;
