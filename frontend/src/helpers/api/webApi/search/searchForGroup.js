import { getOptions } from "../../apiHelper";

const searchGroups = async (query) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group/search?query=${query}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get search results");
	}

	return response.json();
};

export default searchGroups;
