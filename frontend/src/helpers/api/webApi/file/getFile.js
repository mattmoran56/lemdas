import { getOptions } from "../../apiHelper";

const getFile = async (fileId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get file");
	}

	return response.json();
};

export default getFile;
