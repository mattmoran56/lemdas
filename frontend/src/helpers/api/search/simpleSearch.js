import { getOptionsWithBody } from "../apiHelper";

const doSimpleSearch = async (query) => {
	const requestOptions = getOptionsWithBody("POST", JSON.stringify({ query }));

	const response = await fetch(
		`${process.env.REACT_APP_SEARCH_API_URL}/simpleSearch`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get search results");
	}

	return response.json();
};

export default doSimpleSearch;
