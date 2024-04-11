const convertUnixTimestampToStringDate = (timestamp) => {
	const date = new Date(timestamp);
	return date.toLocaleDateString();
};

export default convertUnixTimestampToStringDate;
