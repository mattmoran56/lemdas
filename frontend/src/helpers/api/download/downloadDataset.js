import { getOptions } from "../apiHelper";

const downloadDataset = async (fileId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_UPLOAD_API_URL}/download/dataset/${fileId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to download file");
	}

	return response.blob();
};

export default downloadDataset;
