package model

import "time"

type Asistencia struct {
	TipoAsistencia               string    `json:"tipo_asistencia"`
	CodigoAsistencia             string    `json:"codigo_asistencia"`
	FechaAsistencia              time.Time `json:"fecha_asistencia"`
	NumeroAfiliacion             string    `json:"numero_afiliacion"`
	FinanciadorId                int       `json:"financiador_id"`
	SiglasFinanciador            string    `json:"siglas_financiador"`
	NombreFinanciador            string    `json:"nombre_financiador"`
	CodigoCargoFacturacion       int       `json:"codigo_cargo_facturacion"`
	NombreCargoFacturacion       string    `json:"nombre_cargo_facturacion"`
	TipoAsegurado                string    `json:"tipo_asegurado"`
	MedicoId                     int       `json:"medico_id`
	CodigoMedico                 string    `json:"codigo_medico"`
	SexoMedico                   string    `json:"sexo_medico"`
	PrimerNombreMedico           string    `json:"primer_nombre_medico"`
	SegundoNombreMedico          string    `json:"segundo_nombre_medico"`
	PrimerApellidoMedico         string    `json:"primer_apellido_medico"`
	SegundoApellidoMedico        string    `json:"segundo_apellido_medico"`
	NombreCompletoMmedico        string    `json:"nombre_completo_mmedico"`
	PacienteId                   int       `json:"paciente_id"`
	NombrePaciente               string    `json:"nombre_paciente"`
	ServicioId                   int       `json:"servicio_id"`
	CodigoServicio               string    `json:"codigo_servicio"`
	NombreServicio               string    `json:"nombre_servicio"`
	EspecialidadId               int       `json:"especialidad_id"`
	CodigoEspecialidad           string    `json:"codigo_especialidad"`
	NombreEspecialidad           string    `json:"nombre_especialidad"`
	CambiarServicio              bool      `json:"cambiar_servicio"`
	NumeroExpediente             string    `json:"numero_expediente"`
	CodigoUnidadOrigen           int       `json:"codigo_unidad_origen"`
	NombreUnidadOrigen           string    `json:"nombre_unidad_origen"`
	CodigoAlternoEstablecimiento int       `json:"codigo_alterno_establecimiento"`
	NombreEstablecimiento        string    `json:"nombre_establecimiento"`
}
