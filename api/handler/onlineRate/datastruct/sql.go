package datastruct
const Query1=` SELECT  CONCAT(DATE_FORMAT(pr.operation_time, '%Y-%m-'),FLOOR(DATE_FORMAT(pr.operation_time, '%d') /?)) as time, SUM(pr.duration) as s
								FROM iot_device_process_record pr
								WHERE pr.device_id =?
								AND pr.state = 1 
								AND UNIX_TIMESTAMP(?)<=UNIX_TIMESTAMP(pr.operation_time)
								AND UNIX_TIMESTAMP(pr.operation_time)<UNIX_TIMESTAMP(?)
								GROUP BY CONCAT(DATE_FORMAT(pr.operation_time, '%Y-%m- '),FLOOR(DATE_FORMAT(pr.operation_time, '%d') / 2))`
const Query2=` SELECT  CONCAT(DATE_FORMAT(pr.operation_time, '%Y-%m-%d '),FLOOR(DATE_FORMAT(pr.operation_time, '%H') /?)) as time, SUM(pr.duration) as s
								FROM iot_device_process_record pr
								WHERE pr.device_id =?
								AND pr.state = 1 
								AND UNIX_TIMESTAMP(?)<=UNIX_TIMESTAMP(pr.operation_time)
								AND UNIX_TIMESTAMP(pr.operation_time)<UNIX_TIMESTAMP(?)
								GROUP BY CONCAT(DATE_FORMAT(pr.operation_time, '%Y-%m-%d '),FLOOR(DATE_FORMAT(pr.operation_time, '%H') / 2))`
const Query3 =`SELECT  DATE_FORMAT(pr.operation_time, '%Y-%m-%d %H-%i') as time
								FROM iot_device_process_record pr
								WHERE pr.device_id =?
								AND pr.state = 1 
								AND UNIX_TIMESTAMP(DATE_FORMAT(DATE_SUB(NOW(),INTERVAL ? hour), '%Y-%m-%d %H'))<= UNIX_TIMESTAMP(pr.operation_time)
								AND UNIX_TIMESTAMP(pr.operation_time)<UNIX_TIMESTAMP(DATE_FORMAT(NOW(), '%Y-%m-%d %H'))
								 `