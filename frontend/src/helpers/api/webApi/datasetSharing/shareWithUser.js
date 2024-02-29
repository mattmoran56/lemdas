import { getOptionsWithBody } from "../../apiHelper";

const shareWithUser = async (datasetId, userId, writeAccess) => {
	const requestOptions = getOptionsWithBody("POST", JSON.stringify({ user_id: userId, write_access: writeAccess }));

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset/${datasetId}/share/user`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to share with user");
	}

	return response.json();
};

export default shareWithUser;
