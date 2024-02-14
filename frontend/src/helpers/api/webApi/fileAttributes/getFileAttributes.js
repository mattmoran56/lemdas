import { getOptions } from "../../apiHelper";

const getFileAttributes = async (fileId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}/attribute?orderBy=created_at`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get attributes");
	}

	return response.json();
};

export default getFileAttributes;
