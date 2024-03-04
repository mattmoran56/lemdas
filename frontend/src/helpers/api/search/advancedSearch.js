import { getOptionsWithBody } from "../apiHelper";

const doAdvancedSearch = async (body) => {
	const requestOptions = getOptionsWithBody("POST", JSON.stringify(body));

	const response = await fetch(
		`${process.env.REACT_APP_SEARCH_API_URL}/search`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get search results");
	}

	return response.json();
};

export default doAdvancedSearch;
