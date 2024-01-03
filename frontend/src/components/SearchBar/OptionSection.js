import React from "react";
// TODO: Remove this eslint-disable line
// eslint-disable-next-line import/no-extraneous-dependencies
import { PlusIcon, TrashIcon } from "@heroicons/react/24/outline";

import Button from "../basic/Button";

const OptionSection = ({
	first, lastRemaining, logic, field, term, onChange, onNewOption, onDelete,
}) => {
	return (
		<div className="flex items-center justify-center my-2">
			<div className="w-24">
				{first ? null : (
					<select
						id="logic"
						className="text-black px-3 py-2 rounded-3xl border-2 mx-2"
						value={logic}
						onChange={onChange}
					>
						<option value="and">AND</option>
						<option value="or">OR</option>
					</select>
				)}
			</div>
			{/* TODO: Fit this to metadata (pull from server) */}
			<select
				id="field"
				className="text-black px-3 py-2 rounded-3xl border-2 mx-2"
				value={field}
				onChange={onChange}
			>
				<option value="title">Title</option>
				<option value="description">Description</option>
				<option value="tags">Tags</option>
			</select>
			{/* TODO: Fit this to datatype (e.g. between for dates, contains for text) */}
			<p>Contains</p>
			<input
				className=" bg-transparent border-black border-b-2 mx-4 outline-none"
				placeholder="Search term"
				value={term}
				id="term"
				onChange={onChange}
				autoComplete="off"
			/>
			<Button className="mx-2 rounded-full px-2" onClick={onNewOption}>
				<PlusIcon className="w-4 h-4" />
			</Button>
			{lastRemaining ? <div className="w-4 mx-4" /> : (
				<Button className="mx-2 rounded-full px-2 bg-oxfordblue" onClick={onDelete}>
					<TrashIcon className="w-4 h-4" />
				</Button>
			)}
		</div>
	);
};

export default OptionSection;
