import { getOptionsWithBody } from "../../apiHelper";

const createNewGroup = async (groupName) => {
	const body = JSON.stringify({
		group_name: groupName,
	});
	const requestOptions = getOptionsWithBody("POST", body);
	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to create new group");
	}

	return response.json();
};

export default createNewGroup;
