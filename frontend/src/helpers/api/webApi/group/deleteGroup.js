import { getOptions } from "../../apiHelper";

const deleteGroup = async (groupId) => {
	const requestOptions = getOptions("DELETE");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group/${groupId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to delete group");
	}

	return true;
};

export default deleteGroup;
