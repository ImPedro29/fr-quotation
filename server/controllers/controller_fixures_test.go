package controllers

const returnFromFreteRapidoQuotation = `{"dispatchers":[{"id":"648f3791e585787ed551bbf2","request_id":"648f3791e585787ed551bbf1","registered_number_shipper":"25438296000158","registered_number_dispatcher":"25438296000158","zipcode_origin":29161376,"offers":[{"offer":1,"table_reference":"63b7fd854ed2f3f5dc78b4f5","simulation_type":0,"carrier":{"name":"CORREIOS","registered_number":"34028316000103","state_inscription":"ISENTO","logo":"","reference":281,"company_name":"EMPRESABRASILEIRADECORREIOSETELEGRAFOS"},"service":"Normal","delivery_time":{"days":5,"hours":19,"minutes":34,"estimated_date":"2023-06-23"},"expiration":"2023-07-18T16:57:53.613240369Z","cost_price":78.03,"final_price":0,"weights":{"real":13,"used":17},"correios":{},"original_delivery_time":{"days":5,"hours":19,"minutes":34,"estimated_date":"2023-06-23"}}]}]}`

const mockSendToApp = `{"recipient":{"address":{"zipcode":"01311000"}},"volumes":[{"category":7,"amount":1,"unitary_weight":5,"price":349,"sku":"abc-teste-123","height":0.2,"width":0.2,"length":0.2},{"category":7,"amount":2,"unitary_weight":4,"price":556,"sku":"abc-teste-527","height":0.4,"width":0.6,"length":0.15}]}`
const mockSendToAppInvalid = `{"recipient":{"address":{"zipcode":""}},"volumes":[{"category":7,"amount":1,"unitary_weight":5,"price":349,"sku":"abc-teste-123","height":0.2,"width":0.2,"length":0.2},{"category":7,"amount":2,"unitary_weight":4,"price":556,"sku":"abc-teste-527","height":0.4,"width":0.6,"length":0.15}]}`

const createQuote = `INSERT INTO quotations`
