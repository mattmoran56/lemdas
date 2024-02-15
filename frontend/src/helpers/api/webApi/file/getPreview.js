import { getOptions } from "../../apiHelper";

const getPreviewURL = async (fileId) => {
	const requestOptions = getOptions("GET");
	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}/preview`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get file preview url");
	}

	return response.json();
};

export default getPreviewURL;
