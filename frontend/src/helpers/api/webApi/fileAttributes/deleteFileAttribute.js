import { getOptions } from "../../apiHelper";

const deleteFileAttribute = async (fileId, attributeId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/file/${fileId}/attribute/${attributeId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete attributes");
	}

	return true;
};

export default deleteFileAttribute;
