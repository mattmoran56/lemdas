import { getOptionsWithBody } from "../apiHelper";

const uploadFile = async (files, datasetId, isPublic) => {
	const formData = new FormData();
	for (let i = 0; i < files.length; i += 1) {
		formData.append("file", files[i]);
	}
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
