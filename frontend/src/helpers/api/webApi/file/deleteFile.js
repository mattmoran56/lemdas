import { getOptions } from "../../apiHelper";

const deleteFile = async (fileId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete file");
	}

	return true;
};

export default deleteFile;
