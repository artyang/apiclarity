import moment from "moment";
import COLORS from "utils/scss_variables.module.scss";

export const formatDateBy = (date, format) =>
  !!date ? moment(date).format(format) : "";
export const formatDate = (date) => formatDateBy(date, "MMM Do, YYYY HH:mm:ss");

export const SYSTEM_RISKS = {
  CRITICAL: {
    value: "CRITICAL",
    label: "Critical",
    color: COLORS["color-risk-critical"],
  },
  HIGH: { value: "HIGH", label: "High", color: COLORS["color-risk-high"] },
  MEDIUM: {
    value: "MEDIUM",
    label: "Medium",
    color: COLORS["color-risk-medium"],
  },
  LOW: { value: "LOW", label: "Low", color: COLORS["color-risk-low"] },
  UNKNOWN: { value: "UNKNOWN", label: "Unknown" },
  NO_RISK: { value: "NO_RISK", label: "No known risk" },
  APPROVED: { value: "APPROVED", label: "Approved" },
  NEUTRAL: { value: "NEUTRAL", label: "Neutral" },
  INFO: { value: "INFO", label: "Info", color: COLORS["color-risk-info"] },
};

export const TRACES_TYPES = {
  APIGEE_X: {
    label: "Apigee X",
    value: "APIGEE_X",
    type: "GATEWAY",
    typeLabel: "API Gateway",
  },
  F5_BIG_IP: {
    label: "F5 Big IP",
    value: "F5_BIG_IP",
    type: "LOADBALANCER",
    typeLabel: "Load Balancer",
  },
  KONG_INTERNAL: {
    label: "Kong",
    value: "KONG_INTERNAL",
    type: "GATEWAY",
    typeLabel: "API Gateway",
  },
  TYK_INTERNAL: {
    label: "Tyk",
    value: "TYK_INTERNAL",
    type: "GATEWAY",
    typeLabel: "API Gateway",
  },
};
