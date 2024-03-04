import { getOptionsWithBody } from "../../apiHelper";

const addUserToGroup = async (groupId, userId) => {
	const body = JSON.stringify({
		user_id: userId,
	});
	const requestOptions = getOptionsWithBody("POST", body);
	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group/${groupId}/member`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to add member");
	}

	return response.json();
};

export default addUserToGroup;
