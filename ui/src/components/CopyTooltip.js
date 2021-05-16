import React from 'react';
import Button from '@material-ui/core/Button';
import Tooltip from '@material-ui/core/Tooltip';

export default function CopyTooltip(props) {
  const handleTooltipOpen = () => {
    navigator.clipboard.writeText(props.copy);
  };

  return (
    <Tooltip
      interactive
      title={props.copy}
      placement="right"
    >
      <Button onClick={handleTooltipOpen}>{props.text}</Button>
    </Tooltip>
  );
}

