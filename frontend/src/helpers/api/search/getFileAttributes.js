import { getOptions } from "../apiHelper";

const getFileAttributes = async () => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_SEARCH_API_URL}/fileAttributes`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get attributes");
	}

	return response.json();
};

export default getFileAttributes;
