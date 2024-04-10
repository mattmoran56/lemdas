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

	const data = await response.json();
	return data.attribute_groups;
};

export default getFileAttributes;
