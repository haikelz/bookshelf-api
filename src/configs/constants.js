import dotenv from "dotenv";

dotenv.config();

export const PORT = process.env.PORT;
export const HOST = process.env.HOST;
export const CONDITION = process.env.NODE_ENV;
