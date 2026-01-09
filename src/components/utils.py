import logging
import os
from datetime import datetime, timedelta
from typing import Dict, List, Optional

logger = logging.getLogger(__name__)

def configure_logger(log_level: str = 'INFO') -> None:
    log_level = log_level.upper()
    logging.basicConfig(
        format='%(asctime)s [%(levelname)s] %(name)s: %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        level=getattr(logging, log_level)
    )

def get_env_variables(vars: List[str]) -> Dict[str, str]:
    env_vars = {}
    for var in vars:
        value = os.getenv(var)
        if value is None:
            logger.warning(f"Environment variable '{var}' is not set")
        else:
            env_vars[var] = value
    return env_vars

def parse_datetime(date_str: str, date_format: str = '%Y-%m-%d %H:%M:%S') -> Optional[datetime]:
    try:
        return datetime.strptime(date_str, date_format)
    except ValueError:
        logger.error(f"Failed to parse date string '{date_str}' with format '{date_format}'")
        return None

def add_days(date: datetime, days: int) -> datetime:
    return date + timedelta(days=days)

def is_valid_uuid(uuid_to_test: str) -> bool:
    import re
    regex = re.compile('^[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}$', re.I)
    return bool(regex.match(uuid_to_test))