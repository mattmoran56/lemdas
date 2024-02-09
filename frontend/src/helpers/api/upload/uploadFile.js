import { getOptionsWithBody } from "../apiHelper";

const uploadFile = async (file, datasetId, isPublic) => {
	const formData = new FormData();
	formData.append("file", file);
	formData.append("dataset_id", datasetId);
	formData.append("is_public", isPublic);

	const requestOptions = getOptionsWithBody("POST", formData);

	const response = await fetch(
		`${process.env.REACT_APP_UPLOAD_API_URL}/upload`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to upload file");
	}

	return response.json();
};

export default uploadFile;
