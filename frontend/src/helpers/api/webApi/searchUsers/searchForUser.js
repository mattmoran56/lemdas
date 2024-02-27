import { getOptions } from "../../apiHelper";

const searchUsers = async (query) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/user/search?query=${query}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get search results");
	}

	return response.json();
};

export default searchUsers;
