import logging
import os
from typing import Dict, List, Tuple

# Set up logging configuration
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def load_config(file_path: str) -> Dict:
    try:
        with open(file_path, 'r') as file:
            config = {}
            for line in file:
                key, value = line.strip().split('=')
                config[key] = value
            return config
    except FileNotFoundError:
        logger.error(f"Config file not found: {file_path}")
        return {}

def validate_input(data: Dict) -> bool:
    required_keys = ['name', 'email', 'phone']
    for key in required_keys:
        if key not in data:
            return False
    return True

def process_data(data: List[Dict]) -> Tuple[List, List]:
    successful_records = []
    failed_records = []
    for record in data:
        if validate_input(record):
            successful_records.append(record)
        else:
            failed_records.append(record)
    return successful_records, failed_records

def get_environment_variable(var_name: str) -> str:
    return os.environ.get(var_name, '')